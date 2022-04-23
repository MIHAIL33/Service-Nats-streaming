package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/MIHAIL33/Service-Nats-streaming"
	"github.com/MIHAIL33/Service-Nats-streaming/pkg/cache"
	"github.com/MIHAIL33/Service-Nats-streaming/pkg/handler"
	"github.com/MIHAIL33/Service-Nats-streaming/pkg/repository"
	"github.com/MIHAIL33/Service-Nats-streaming/pkg/service"
	"github.com/MIHAIL33/Service-Nats-streaming/pkg/stream"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title L0
// @version 1.0
// @description L0 task, service with Nats-streaming

// @host localhost:8000
// @BasePath /

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize postgres: %s", err.Error())
	}

	sc, _ := stan.Connect(viper.GetString("nats.clusterID"), viper.GetString("nats.clientSubscriber"), stan.NatsURL(viper.GetString("nats.serverURL")))
	defer sc.Close()

	repos := repository.NewRepository(db)
	
	cache := cache.NewCache()

	services := service.NewService(repos, cache)

	err = services.AddAllInCache()
	if err != nil {
		logrus.Fatalf("failed to load cache: %s", err.Error())
	}

	handlers := handler.NewHandler(services)
	
	stream.NewStream(sc, services)

	srv := new(app.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Println("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	logrus.Println("App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
