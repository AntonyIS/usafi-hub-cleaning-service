package cmd

import (
	"github.com/AntonyIS/usafi-hub-cleaning-service/config"
	"github.com/AntonyIS/usafi-hub-cleaning-service/internal/adapter/app"
	"github.com/AntonyIS/usafi-hub-cleaning-service/internal/adapter/logger"
	"github.com/AntonyIS/usafi-hub-cleaning-service/internal/adapter/repository"
	"github.com/AntonyIS/usafi-hub-cleaning-service/internal/core/services"
)

func RunService() {
	logger, err := logger.NewDefaultLogger()
	if err != nil {
		panic(err)
	}

	config, err := config.NewConfig(logger)
	if err != nil {
		panic(err)
	}

	serviceRepo, _ := repository.NewServicePostgresClient(*config)
	requestRepo, _ := repository.NewRequestPostgresClient(*config)
	reviewRepo, _ := repository.NewReviewPostgresClient(*config)

	serviceService := services.NewServiceServiceManagement(serviceRepo, logger)
	requestService := services.NewRequestServiceManagement(requestRepo, logger)
	reviewService := services.NewReviewServiceManagement(reviewRepo, logger)

	app.InitGinRoutes(serviceService, requestService, reviewService, *config, logger)
}
