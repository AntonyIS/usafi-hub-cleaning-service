package app

import (
	"net/http"

	"github.com/AntonyIS/usafi-hub-cleaning-service/internal/core/domain"
	"github.com/AntonyIS/usafi-hub-cleaning-service/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type GinHandler interface {
	Home(ctx *gin.Context)
	Healthcheck(ctx *gin.Context)
	CreateService(ctx *gin.Context)
	GetServiceById(ctx *gin.Context)
	GetServices(ctx *gin.Context)
	UpdateService(ctx *gin.Context)
	DeleteService(ctx *gin.Context)
	CreateRequest(ctx *gin.Context)
	GetRequestById(ctx *gin.Context)
	GetRequests(ctx *gin.Context)
	UpdateRequest(ctx *gin.Context)
	DeleteRequest(ctx *gin.Context)
	AssignCleaner(ctx *gin.Context)
	GetRequestByClient(ctx *gin.Context)
	GetRequestByCleaner(ctx *gin.Context)
	CreateReview(ctx *gin.Context)
	GetReviewById(ctx *gin.Context)
	UpdateReview(ctx *gin.Context)
	DeleteReview(ctx *gin.Context)
	GetReviewByClient(ctx *gin.Context)
	GetReviewByCleaner(ctx *gin.Context)
}

type handler struct {
	serviceService ports.ServiceService
	requestService ports.RequestService
	reviewService  ports.ReviewService
}

func NewGinHandler(serviceService ports.ServiceService, requestService ports.RequestService, reviewService ports.ReviewService) GinHandler {
	routerHandler := handler{
		serviceService: serviceService,
		requestService: requestService,
		reviewService:  reviewService,
	}
	return routerHandler
}

func (h handler) Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"responseMessage": "Usafihub Cleaning Service",
		"responseCode":    http.StatusOK,
	})
}

func (h handler) Healthcheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"responseMessage": "Usafihub Cleaning Service Health Check",
		"responseCode":    http.StatusOK,
	})
}

func (h handler) CreateService(ctx *gin.Context) {
	var service domain.Service
	if err := ctx.ShouldBindJSON(&service); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusBadRequest,
		})
		return
	}

	dbService, err := h.serviceService.CreateService(service)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"responseMessage": "Service created successfully",
		"responseCode":    http.StatusCreated,
		"data":            dbService,
	})
}

func (h handler) GetServiceById(ctx *gin.Context) {
	serviceId := ctx.Param("service_id")

	service, err := h.serviceService.GetServiceById(serviceId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseMessage": "Service found",
		"responseCode":    http.StatusOK,
		"data":            service,
	})
}

func (h handler) GetServices(ctx *gin.Context) {
	services, err := h.serviceService.GetServices()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseMessage": "Services found",
		"responseCode":    http.StatusOK,
		"data":            services,
		"responsecount":           len(*services),
	})
}

func (h handler) UpdateService(ctx *gin.Context) {
	var service domain.Service
	if err := ctx.ShouldBindJSON(&service); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusBadRequest,
		})
		return
	}

	updatedService, err := h.serviceService.UpdateService(service)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseMessage": "Service updated successfully",
		"responseCode":    http.StatusOK,
		"data":            updatedService,
	})
}

func (h handler) DeleteService(ctx *gin.Context) {
	serviceId := ctx.Param("service_id")

	err := h.serviceService.DeleteService(serviceId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseMessage": "Service deleted successfully",
		"responseCode":    http.StatusOK,
	})
}

func (h handler) CreateRequest(ctx *gin.Context) {
	var request domain.Request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusBadRequest,
		})
		return
	}

	dbRequest, err := h.requestService.CreateRequest(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"responseMessage": "Request created successfully",
		"responseCode":    http.StatusCreated,
		"data":            dbRequest,
	})
}

func (h handler) GetRequestById(ctx *gin.Context) {
	requestId := ctx.Param("request_id")

	request, err := h.requestService.GetRequestById(requestId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseMessage": "Request found",
		"responseCode":    http.StatusOK,
		"data":            request,
	})
}

func (h handler) GetRequests(ctx *gin.Context) {
	requests, err := h.requestService.GetRequests()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseMessage": "Requests found",
		"responseCode":    http.StatusOK,
		"data":            requests,
	})
}

func (h handler) UpdateRequest(ctx *gin.Context) {
	var request domain.Request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusBadRequest,
		})
		return
	}

	updatedRequest, err := h.requestService.UpdateRequest(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseMessage": "Request updated successfully",
		"responseCode":    http.StatusOK,
		"data":            updatedRequest,
	})
}

func (h handler) DeleteRequest(ctx *gin.Context) {
	requestId := ctx.Param("request_id")

	err := h.requestService.DeleteRequest(requestId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseMessage": "Request deleted successfully",
		"responseCode":    http.StatusOK,
	})
}

func (h handler) AssignCleaner(ctx *gin.Context) {
	requestId := ctx.Param("request_id")
	cleanerId := ctx.Param("cleaner_id")

	err := h.requestService.AssignCleaner(requestId, cleanerId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseMessage": "Cleaner assigned successfully",
		"responseCode":    http.StatusOK,
	})
}

func (h handler) GetRequestByClient(ctx *gin.Context) {
	clientId := ctx.Param("client_id")

	requests, err := h.requestService.GetRequestById(clientId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseMessage": "Requests found for client",
		"responseCode":    http.StatusOK,
		"data":            requests,
	})
}

func (h handler) GetRequestByCleaner(ctx *gin.Context) {
	cleanerId := ctx.Param("cleaner_id")

	requests, err := h.requestService.GetRequestByCleaner(cleanerId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseMessage": "Requests found for cleaner",
		"responseCode":    http.StatusOK,
		"data":            requests,
	})
}

func (h handler) CreateReview(ctx *gin.Context) {
	var review domain.Reviews
	if err := ctx.ShouldBindJSON(&review); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusBadRequest,
		})
		return
	}

	dbReview, err := h.reviewService.CreateReview(review)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"responseMessage": "Review created successfully",
		"responseCode":    http.StatusCreated,
		"data":            dbReview,
	})
}

func (h handler) GetReviewById(ctx *gin.Context) {
	reviewId := ctx.Param("review_id")

	review, err := h.reviewService.GetReviewById(reviewId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseMessage": "Review found",
		"responseCode":    http.StatusOK,
		"data":            review,
	})
}

func (h handler) UpdateReview(ctx *gin.Context) {
	var review domain.Reviews
	if err := ctx.ShouldBindJSON(&review); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusBadRequest,
		})
		return
	}

	updatedReview, err := h.reviewService.UpdateReview(review)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseMessage": "Review updated successfully",
		"responseCode":    http.StatusOK,
		"data":            updatedReview,
	})
}

func (h handler) DeleteReview(ctx *gin.Context) {
	reviewId := ctx.Param("review_id")

	err := h.reviewService.DeleteReview(reviewId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseMessage": "Review deleted successfully",
		"responseCode":    http.StatusOK,
	})
}

func (h handler) GetReviewByClient(ctx *gin.Context) {
	clientId := ctx.Param("client_id")

	reviews, err := h.reviewService.GetReviewByClient(clientId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseMessage": "Reviews found for client",
		"responseCode":    http.StatusOK,
		"data":            reviews,
	})
}

func (h handler) GetReviewByCleaner(ctx *gin.Context) {
	cleanerId := ctx.Param("cleaner_id")

	reviews, err := h.reviewService.GetReviewByCleaner(cleanerId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
			"responseCode":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseMessage": "Reviews found for cleaner",
		"responseCode":    http.StatusOK,
		"data":            reviews,
	})
}
