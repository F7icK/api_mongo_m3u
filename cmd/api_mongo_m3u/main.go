package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/server"
	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/server/handlers"
	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/service"
	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/service/workers"
	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/types/config"
	"github.com/F7icK/api_mongo_m3u/internal/clients/repository"
	"github.com/F7icK/api_mongo_m3u/pkg/database/mongo"
	"github.com/F7icK/api_mongo_m3u/pkg/logger"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func main() {
	configPath := new(string)

	flag.StringVar(configPath, "config-path", "./config/config-local.yaml", "specify path to yaml")
	flag.Parse()

	configFile, err := os.Open(*configPath)
	if err != nil {
		logger.LogFatal(errors.Wrap(err, "err with os.Open config"))
		return
	}

	cfg := config.Config{}
	if err = yaml.NewDecoder(configFile).Decode(&cfg); err != nil {
		logger.LogFatal(errors.Wrap(err, "err with Decode config"))
		return
	}

	if err = logger.NewLogger(cfg.Telegram); err != nil {
		logger.LogFatal(err)
		return
	}

	mongoDB, err := mongo.NewMongoDB(cfg.MongoDB)
	if err != nil {
		logger.LogFatal(errors.Wrap(err, "err with NewMongoDB"))
		return
	}

	repos := repository.NewRepository(mongoDB.Client)

	if err != nil {
		logger.LogFatal(errors.Wrap(err, "err with EmailService"))
	}

	worker := workers.New()

	services := service.NewService(
		&cfg,
		repos,
		worker,
	)

	endpoints := handlers.NewHandlers(services, &cfg)

	if err = worker.StartWorker(); err != nil {
		logger.LogFatal(errors.Wrap(err, "err with StartWorker"))
	}

	srv := server.NewServer(&cfg, endpoints)

	go func() {
		if err = srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.LogFatal(errors.Wrap(err, "err with NewRouter"))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	if err = worker.StopWorker(); err != nil {
		logger.LogFatal(errors.Wrap(err, "err with StopWorker"))
	}

	if err = srv.Shutdown(context.Background()); err != nil {
		logger.LogFatal(errors.Wrap(err, "failed to stop server"))
	}
}
