package main

import "github.com/sirupsen/logrus"
import "go.uber.org/zap"

func main() {
	log := logrus.New()

	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	log.Info("This is an info message")
	log.Warn("This is a warn message")
	log.Error("This is an error message")

	log.WithFields(logrus.Fields{
		"username": "nv",
		"method":   "GET",
	}).Info("User logged in")

	// zap
	logger, err := zap.NewProduction()
	if err != nil {
		log.Println("Error initializing zap logger")
	}

	defer logger.Sync()

	logger.Info("This is an info message from zap")
	logger.Info("User logged in", zap.String("username", "nv"), zap.String("method", "POST"))
}
