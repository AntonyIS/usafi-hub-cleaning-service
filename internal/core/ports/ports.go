package ports

import "github.com/AntonyIS/usafi-hub-cleaning-service/internal/core/domain"

type ServiceService interface {
	CreateService(service domain.Service) (*domain.Service, error)
	GetServiceById(service_id string) (*domain.Service, error)
	GetServices() (*[]domain.Service, error)
	UpdateService(service domain.Service) (*domain.Service, error)
	DeleteService(service_id string) error
}

type RequestService interface {
	CreateRequest(request domain.Request) (*domain.Request, error)
	GetRequestById(request_id string) (*domain.Request, error)
	GetRequests() (*[]domain.Request, error)
	UpdateRequest(request domain.Request) (*domain.Request, error)
	DeleteRequest(request_id string) error
	AssignCleaner(request_id, cleaner_id string) error
	GetRequestByClient(client_id string) error
	GetRequestByCleaner(cleaner_id string) error
}

type ReviewService interface {
	CreateReview(review domain.Reviews) (*domain.Reviews, error)
	GetReviewById(review_id string) (*domain.Reviews, error)
	UpdateReview(review domain.Reviews) (*domain.Reviews, error)
	DeleteReview(review_id string) error
	GetReviewByClient(client_id string) error
	GetReviewByCleaner(cleaner_id string) error
}

type ServiceRepository interface {
	CreateService(service domain.Service) (*domain.Service, error)
	GetServiceById(service_id string) (*domain.Service, error)
	GetServices() (*[]domain.Service, error)
	UpdateService(service domain.Service) (*domain.Service, error)
	DeleteService(service_id string) error
}

type RequestRepository interface {
	CreateRequest(request domain.Request) (*domain.Request, error)
	GetRequestById(request_id string) (*domain.Request, error)
	GetRequests() (*[]domain.Request, error)
	UpdateRequest(request domain.Request) (*domain.Request, error)
	DeleteRequest(request_id string) error
	AssignCleaner(request_id, cleaner_id string) error
	GetRequestByClient(client_id string) error
	GetRequestByCleaner(cleaner_id string) error
}

type ReviewRepository interface {
	CreateReview(review domain.Reviews) (*domain.Reviews, error)
	GetReviewById(review_id string) (*domain.Reviews, error)
	UpdateReview(review domain.Reviews) (*domain.Reviews, error)
	DeleteReview(review_id string) error
	GetReviewByClient(client_id string) error
	GetReviewByCleaner(cleaner_id string) error
}

type LoggerService interface {
	Info(message string)
	Warning(message string)
	Error(message string)
}
