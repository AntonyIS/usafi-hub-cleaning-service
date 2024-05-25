package app

import (
	"fmt"
	"log"

	"github.com/AntonyIS/usafi-hub-cleaning-service/config"
	"github.com/AntonyIS/usafi-hub-cleaning-service/internal/core/ports"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitGinRoutes(serviceService ports.ServiceService, requestService ports.RequestService, reviewService ports.ReviewService, config config.Config, logger ports.LoggerService) {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	handler := NewGinHandler(
		serviceService,
		requestService,
		reviewService,
	)

	// Define routes
	homeRoutes := router.Group("/")
	servicesRoutes := router.Group("/services/v1")
	requestsRoutes := router.Group("/requests/v1")
	reviewsRoutes := router.Group("/reviews/v1")

	// middleware := NewMiddleware(logger, config.SECRET_KEY)

	// servicesRoutes.Use(middleware.GinAuthMiddleware())
	// requestsRoutes.Use(middleware.GinAuthMiddleware())
	// reviewsRoutes.Use(middleware.GinAuthMiddleware())
	// homeRoutes.Use(middleware.GinAuthMiddleware())

	// Home routes
	homeRoutes.GET("/", handler.Home)
	homeRoutes.GET("/health-check", handler.Healthcheck)

	// Services routes
	servicesRoutes.POST("/", handler.CreateService)
	servicesRoutes.GET("/:service_id", handler.GetServiceById)
	servicesRoutes.GET("/", handler.GetServices)
	servicesRoutes.PUT("/:service_id", handler.UpdateService)
	servicesRoutes.DELETE("/:service_id", handler.DeleteService)

	// Requests routes
	requestsRoutes.POST("/", handler.CreateRequest)
	requestsRoutes.GET("/:request_id", handler.GetRequestById)
	requestsRoutes.GET("/", handler.GetRequests)
	requestsRoutes.PUT("/:request_id", handler.UpdateRequest)
	requestsRoutes.DELETE("/:request_id", handler.DeleteRequest)
	requestsRoutes.POST("/:request_id/assign-cleaner/:cleaner_id", handler.AssignCleaner)
	requestsRoutes.GET("/client/:client_id", handler.GetRequestByClient)
	requestsRoutes.GET("/cleaner/:cleaner_id", handler.GetRequestByCleaner)

	// Reviews routes
	reviewsRoutes.POST("/", handler.CreateReview)
	reviewsRoutes.GET("/:review_id", handler.GetReviewById)
	reviewsRoutes.PUT("/:review_id", handler.UpdateReview)
	reviewsRoutes.DELETE("/:review_id", handler.DeleteReview)
	reviewsRoutes.GET("/client/:client_id", handler.GetReviewByClient)
	reviewsRoutes.GET("/cleaner/:cleaner_id", handler.GetReviewByCleaner)

	log.Printf("Server running on port 0.0.0.0:%s", config.SERVER_PORT)
	router.Run(fmt.Sprintf(":%s", config.SERVER_PORT))
}
