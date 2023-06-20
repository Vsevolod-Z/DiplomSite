package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"diplomsite/internal/handler"
	"diplomsite/server"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	initLogrus()

	if err := initConfig(); err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "main",
			"function": "initConfig()",
			"error":    err,
			"config":   viper.AllKeys(),
		}).Fatal("Error in reading the configuration")

	}
	handlers := handler.NewHandler()
	srv := new(server.Server)
	go func() {
		if err := srv.Start(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.WithFields(logrus.Fields{
				"package":  "main",
				"function": "server.Start()",
				"error":    err,
				"config":   viper.AllKeys(),
				"port":     viper.GetString("port"),
			}).Fatal("Server startup error")

		}
		logrus.WithFields(logrus.Fields{
			"package":  "main",
			"function": "main()",
			"config":   viper.AllKeys(),
			"port":     viper.GetString("port"),
		}).Info("The server is running correctly ")
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("SERVER STOP!")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

}

func initConfig() error {
	viper.SetConfigType("yml")
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
func initLogrus() {
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(file)
}
