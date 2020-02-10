package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"simple-http/internal/configs"
	"simple-http/internal/domain"
	"simple-http/internal/infra/http"
)

func main()  {
	config := &configs.Config{Port: ":3000", LogLevel: "info"}
	//config, err := configs.Parse()
	//if err != nil {
	//	if err, ok := err.(*flags.Error); ok {
	//		fmt.Println(err)
	//		os.Exit(0)
	//	}
	//
	//	fmt.Printf("Invalid args: %v\n", err)
	//	os.Exit(1)
	//}

	logger := logrus.New()

	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger.SetLevel(level)

	service := domain.NewService(logger)


	httpAdapter, err := http.NewAdapter(logger, config, service)

	if err != nil {
		logger.WithError(err).Fatal("Error creating new HTTP adapter!")
	}

	if err := httpAdapter.ListenAndServe(); err != nil {
		logger.WithError(err).Fatal("Error listening to server")
	}
}
