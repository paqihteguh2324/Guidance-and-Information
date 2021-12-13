package main

import (
	"context"
	"database/sql"
	_ "expvar"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"sadhelx-be-guidelines/service"
	"sadhelx-be-guidelines/service/repository"

	"sadhelx-be-guidelines/util"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "usermgmt",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	loggingMiddleware := service.LoggingMiddleware(logger)

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	configs := util.NewConfigurations(logger)
	initDb(configs)
	defer db.Close()

	var httpAddr = flag.String("http", ":"+configs.ServerPort, "http listen address")

	flag.Parse()
	ctx := context.Background()

	var srv service.Service
	{
		svcRepository := repository.NewRepo(db, logger)

		srv = service.NewService(
			svcRepository,
			// configs,
			logger,
		)
	}
	endpoints := service.MakeGuidelinesEndpoints(srv)

	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("listening-on", configs.ServerPort)
		handler := service.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, loggingMiddleware(handler))

	}()

	level.Error(logger).Log("exit", <-errChan)

}

func initDb(confs *util.Configurations) {
	// config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		confs.DBHost, confs.DBPort,
		confs.DBUser, confs.DBPass, confs.DBName)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

}
