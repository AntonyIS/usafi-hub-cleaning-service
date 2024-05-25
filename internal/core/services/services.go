package services

import (
	"time"

	"github.com/AntonyIS/usafi-hub-cleaning-service/internal/core/domain"
	"github.com/AntonyIS/usafi-hub-cleaning-service/internal/core/ports"
	"github.com/google/uuid"
)

type ServiceServiceManagement struct {
	repo   ports.ServiceRepository
	logger ports.LoggerService
}

type RequestServiceManagement struct {
	repo   ports.RequestRepository
	logger ports.LoggerService
}

type ReviewServiceManagement struct {
	repo   ports.ReviewRepository
	logger ports.LoggerService
}

func NewServiceServiceManagement(repo ports.ServiceRepository, logger ports.LoggerService) *ServiceServiceManagement {
	service := ServiceServiceManagement{
		repo:   repo,
		logger: logger,
	}

	return &service
}

func NewRequestServiceManagement(repo ports.RequestRepository, logger ports.LoggerService) *RequestServiceManagement {
	service := RequestServiceManagement{
		repo:   repo,
		logger: logger,
	}
	return &service
}

func NewReviewServiceManagement(repo ports.ReviewRepository, logger ports.LoggerService) *ReviewServiceManagement {
	service := ReviewServiceManagement{
		repo:   repo,
		logger: logger,
	}
	return &service
}

// Service Methods
func (svc ServiceServiceManagement) CreateService(service domain.Service) (*domain.Service, error) {
	service.ServiceId = uuid.New().String()
	service.CreatedAt = time.Now()
	service.UpdatedAt = time.Now()
	return svc.repo.CreateService(service)
}

func (svc ServiceServiceManagement) GetServiceById(service_id string) (*domain.Service, error) {
	return svc.repo.GetServiceById(service_id)
}

func (svc ServiceServiceManagement) GetServices() (*[]domain.Service, error) {
	return svc.repo.GetServices()
}

func (svc ServiceServiceManagement) UpdateService(service domain.Service) (*domain.Service, error) {
	service.UpdatedAt = time.Now()
	return svc.repo.UpdateService(service)
}

func (svc ServiceServiceManagement) DeleteService(service_id string) error {
	return svc.repo.DeleteService(service_id)
}

// Request Methods
func (svc RequestServiceManagement) CreateRequest(request domain.Request) (*domain.Request, error) {
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	return svc.repo.CreateRequest(request)
}

func (svc RequestServiceManagement) GetRequestById(request_id string) (*domain.Request, error) {
	return svc.repo.GetRequestById(request_id)
}

func (svc RequestServiceManagement) GetRequests() (*[]domain.Request, error) {
	return svc.repo.GetRequests()
}

func (svc RequestServiceManagement) UpdateRequest(request domain.Request) (*domain.Request, error) {
	request.UpdatedAt = time.Now()
	return svc.repo.UpdateRequest(request)
}

func (svc RequestServiceManagement) DeleteRequest(request_id string) error {
	return svc.repo.DeleteRequest(request_id)
}

func (svc RequestServiceManagement) AssignCleaner(request_id, cleaner_id string) error {
	return svc.repo.AssignCleaner(request_id, cleaner_id)
}

func (svc RequestServiceManagement) GetRequestByClient(client_id string) (*[]domain.Request, error) {
	return svc.repo.GetRequestByClient(client_id)
}

func (svc RequestServiceManagement) GetRequestByCleaner(cleaner_id string) (*[]domain.Request, error) {
	return svc.repo.GetRequestByCleaner(cleaner_id)
}

// Review Methods
func (svc ReviewServiceManagement) CreateReview(review domain.Reviews) (*domain.Reviews, error) {
	review.CreatedAt = time.Now()
	review.UpdatedAt = time.Now()
	return svc.repo.CreateReview(review)
}

func (svc ReviewServiceManagement) GetReviewById(review_id string) (*domain.Reviews, error) {
	return svc.repo.GetReviewById(review_id)
}

func (svc ReviewServiceManagement) UpdateReview(review domain.Reviews) (*domain.Reviews, error) {
	review.UpdatedAt = time.Now()
	return svc.repo.UpdateReview(review)
}

func (svc ReviewServiceManagement) DeleteReview(review_id string) error {
	return svc.repo.DeleteReview(review_id)
}

func (svc ReviewServiceManagement) GetReviewByClient(client_id string) (*[]domain.Reviews, error) {
	return svc.repo.GetReviewByClient(client_id)
}

func (svc ReviewServiceManagement) GetReviewByCleaner(cleaner_id string) (*[]domain.Reviews, error) {
	return svc.repo.GetReviewByCleaner(cleaner_id)
}
