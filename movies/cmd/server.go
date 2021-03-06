package main

import (
	"context"
	"fmt"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/jackc/pgx"
	"github.com/julienschmidt/httprouter"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
	"micromovies2/movies"
	"micromovies2/movies/pb"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var pool *pgx.ConnPool

func main() {
	//zerolog
	logger := zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
	var (
		console  bool
		httpAddr string
		gRPCAddr string
	)
	flag.StringVarP(&httpAddr, "http", "H", ":8082", "http listen address")
	flag.StringVarP(&gRPCAddr, "grpc", "g", ":8081", "GRPC Address")
	flag.BoolVarP(&console, "console", "c", false, "turns on pretty console logging")
	flag.Parse()
	logger.Info().Msg("starting grpc server at" + string(gRPCAddr))
	if console {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	ctx := context.Background()
	//database
	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "127.0.0.1",
			Port:     26257,
			User:     "app_user",
			Database: "app_database",
			//Logger: logger, todo: fix logger
		},
		MaxConnections: 5,
	}
	pool, err := pgx.NewConnPool(connPoolConfig)
	if err != nil {
		logger.Fatal().Err(err).Msg("Unable to create connection pool")
	}
	/*db, err := sql.Open("postgres", "postgresql://app_user@localhost:26257/app_database?sslmode=disable")
	if err != nil {
		logger.Fatal().Err(err).Msg("db connection failed")
	}*/

	//instrumentation
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "movies_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "movies_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	// init movies service
	var svc movies.Service
	svc = movies.NewService(pool, logger)
	//wire logging
	svc = movies.LoggingMiddleware{logger, svc}
	//wire instrumentation
	svc = movies.InstrumentingMiddleware{requestCount, requestLatency, svc}
	errChan := make(chan error)
	// creating Endpoints struct
	endpoints := movies.Endpoints{
		GetMoviesEndpoint:    movies.MakeGetMoviesEndpoint(svc),
		GetMovieByIdEndpoint: movies.MakeGetMovieByIdEndpoint(svc),
		NewMovieEndpoint:     movies.MakeNewMovieEndpoint(svc),
		DeleteMovieEndpoint:  movies.MakeDeleteMovieEndpoint(svc),
		UpdateMovieEndpoint:  movies.MakeUpdateMovieEndpoint(svc),
	}
	//execute grpc server
	go func() {
		listener, err := net.Listen("tcp", gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}
		handler := movies.NewGRPCServer(ctx, endpoints)
		grpcServer := grpc.NewServer()
		pb.RegisterMoviesServer(grpcServer, handler)
		errChan <- grpcServer.Serve(listener)
	}()
	// HTTP transport
	go func() {
		//httprouter initialization
		router := httprouter.New()
		//handler will be used for net/http handle compatibility
		router.Handler("GET", "/metrics", promhttp.Handler())
		errChan <- http.ListenAndServe(httpAddr, router)
	}()
	//Handle os signals
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	logger.Error().Err(<-errChan).Msg("")
}
