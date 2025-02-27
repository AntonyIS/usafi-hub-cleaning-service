package repository

import (
	"database/sql"
	"fmt"

	"github.com/AntonyIS/usafi-hub-cleaning-service/config"
	"github.com/AntonyIS/usafi-hub-cleaning-service/internal/core/domain"
	_ "github.com/lib/pq"
)

type postgresClient struct {
	db               *sql.DB
	serviceTablename string
	requestablename  string
	reviewTablename  string
}

func NewServicePostgresClient(config config.Config) (*postgresClient, error) {
	dbname := config.POSTGRES_DB
	user := config.POSTGRES_USER
	password := config.POSTGRES_PASSWORD
	port := config.POSTGRES_PORT
	host := config.POSTGRES_HOST

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	queryString := fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %s (
        service_id VARCHAR(255) PRIMARY KEY UNIQUE,
        name VARCHAR(255) NOT NULL,
        description TEXT,
        price_per_hour FLOAT NOT NULL,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    )
	`, config.SERVICE_TABLE)

	_, err = db.Exec(queryString)
	if err != nil {
		return nil, err
	}
	return &postgresClient{
		db:               db,
		serviceTablename: config.SERVICE_TABLE,
		requestablename:  config.REQUEST_TABLE,
		reviewTablename:  config.REVIEWS_TABLE,
	}, nil
}

func NewRequestPostgresClient(config config.Config) (*postgresClient, error) {
	dbname := config.POSTGRES_DB
	user := config.POSTGRES_USER
	password := config.POSTGRES_PASSWORD
	port := config.POSTGRES_PORT
	host := config.POSTGRES_HOST

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	queryString := fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %s (
        request_id VARCHAR(255) PRIMARY KEY UNIQUE,
        client_id VARCHAR(255) NOT NULL,
        cleaner_id VARCHAR(255) NOT NULL,
        service_id VARCHAR(255) NOT NULL,
        requested_date TIMESTAMP NOT NULL,
        status VARCHAR(50) NOT NULL,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    )
	`, config.REQUEST_TABLE)

	_, err = db.Exec(queryString)
	if err != nil {
		return nil, err
	}
	return &postgresClient{
		db:               db,
		serviceTablename: config.SERVICE_TABLE,
		requestablename:  config.REQUEST_TABLE,
		reviewTablename:  config.REVIEWS_TABLE,
	}, nil
}

func NewReviewPostgresClient(config config.Config) (*postgresClient, error) {
	dbname := config.POSTGRES_DB
	user := config.POSTGRES_USER
	password := config.POSTGRES_PASSWORD
	port := config.POSTGRES_PORT
	host := config.POSTGRES_HOST

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	queryString := fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %s (
        review_id VARCHAR(255) PRIMARY KEY UNIQUE,
        request_id VARCHAR(255) NOT NULL,
        client_id VARCHAR(255) NOT NULL,
        cleaner_id VARCHAR(255) NOT NULL,
        rating VARCHAR(50) NOT NULL,
        comment TEXT,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    )
	`, config.REVIEWS_TABLE)

	_, err = db.Exec(queryString)
	if err != nil {
		return nil, err
	}
	return &postgresClient{
		db:               db,
		serviceTablename: config.SERVICE_TABLE,
		requestablename:  config.REQUEST_TABLE,
		reviewTablename:  config.REVIEWS_TABLE,
	}, nil
}

// CreateService creates a service  using
func (svc postgresClient) CreateService(service domain.Service) (*domain.Service, error) {
	query := fmt.Sprintf(`
        INSERT INTO %s (service_id, name, description, price_per_hour, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
    `, svc.serviceTablename)

	_, err := svc.db.Exec(query,
		service.ServiceId,
		service.Name,
		service.Description,
		service.PricePerHour,
		service.CreatedAt,
		service.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return svc.GetServiceById(service.ServiceId)
}

// GetServiceById retrieves a service  using service id from the services table
func (svc postgresClient) GetServiceById(serviceId string) (*domain.Service, error) {
	query := fmt.Sprintf(`
        SELECT service_id, name, description, price_per_hour, created_at, updated_at
        FROM %s
        WHERE service_id = $1
    `, svc.serviceTablename)

	var service domain.Service
	err := svc.db.QueryRow(query, serviceId).Scan(
		&service.ServiceId,
		&service.Name,
		&service.Description,
		&service.PricePerHour,
		&service.CreatedAt,
		&service.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &service, nil
}

// GetServices retrieves all services from the services table
func (svc postgresClient) GetServices() (*[]domain.Service, error) {
	query := fmt.Sprintf(`
        SELECT service_id, name, description, price_per_hour, created_at, updated_at
        FROM %s
    `, svc.serviceTablename)

	rows, err := svc.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []domain.Service
	for rows.Next() {
		var service domain.Service
		err := rows.Scan(
			&service.ServiceId,
			&service.Name,
			&service.Description,
			&service.PricePerHour,
			&service.CreatedAt,
			&service.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &services, nil
}

// UpdateService updates an existing service in the services table
func (svc postgresClient) UpdateService(service domain.Service) (*domain.Service, error) {
	query := fmt.Sprintf(`
        UPDATE %s
        SET name = $2, description = $3, price_per_hour = $4, updated_at = $5
        WHERE service_id = $1
    `, svc.serviceTablename)

	_, err := svc.db.Exec(query,
		service.ServiceId,
		service.Name,
		service.Description,
		service.PricePerHour,
		service.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return svc.GetServiceById(service.ServiceId)
}

// DeleteService deletes a service from the services table
func (svc postgresClient) DeleteService(serviceId string) error {
	query := fmt.Sprintf(`
        DELETE FROM %s
        WHERE service_id = $1
    `, svc.serviceTablename)

	_, err := svc.db.Exec(query, serviceId)
	if err != nil {
		return err
	}
	return nil
}

func (svc postgresClient) CreateRequest(request domain.Request) (*domain.Request, error) {
	query := fmt.Sprintf(`
        INSERT INTO %s (request_id, client_id, cleaner_id, service_id, requested_date, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `, svc.requestablename)

	_, err := svc.db.Exec(query,
		request.RequestId,
		request.ClientId,
		request.CleanerId,
		request.ServiceId,
		request.RequestedDate,
		request.Status,
		request.CreatedAt,
		request.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return svc.GetRequestById(request.RequestId)
}

func (svc postgresClient) GetRequestById(requestId string) (*domain.Request, error) {
	query := fmt.Sprintf(`
        SELECT request_id, client_id, cleaner_id, service_id, requested_date, status, created_at, updated_at
        FROM %s
        WHERE request_id = $1
    `, svc.requestablename)

	var request domain.Request
	err := svc.db.QueryRow(query, requestId).Scan(
		&request.RequestId,
		&request.ClientId,
		&request.CleanerId,
		&request.ServiceId,
		&request.RequestedDate,
		&request.Status,
		&request.CreatedAt,
		&request.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &request, nil
}

func (svc postgresClient) GetRequests() (*[]domain.Request, error) {
	query := fmt.Sprintf(`
        SELECT request_id, client_id, cleaner_id, service_id, requested_date, status, created_at, updated_at
        FROM %s
    `, svc.requestablename)

	rows, err := svc.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []domain.Request
	for rows.Next() {
		var request domain.Request
		err := rows.Scan(
			&request.RequestId,
			&request.ClientId,
			&request.CleanerId,
			&request.ServiceId,
			&request.RequestedDate,
			&request.Status,
			&request.CreatedAt,
			&request.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &requests, nil
}

func (svc postgresClient) UpdateRequest(request domain.Request) (*domain.Request, error) {
	query := fmt.Sprintf(`
        UPDATE %s
        SET client_id = $2, cleaner_id = $3, service_id = $4, requested_date = $5, status = $6, updated_at = $7
        WHERE request_id = $1
    `, svc.requestablename)

	_, err := svc.db.Exec(query,
		request.RequestId,
		request.ClientId,
		request.CleanerId,
		request.ServiceId,
		request.RequestedDate,
		request.Status,
		request.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return svc.GetRequestById(request.RequestId)
}

func (svc postgresClient) DeleteRequest(requestId string) error {
	query := fmt.Sprintf(`
        DELETE FROM %s
        WHERE request_id = $1
    `, svc.requestablename)

	_, err := svc.db.Exec(query, requestId)
	if err != nil {
		return err
	}
	return nil
}

func (svc postgresClient) AssignCleaner(requestId, cleanerId string) error {
	query := fmt.Sprintf(`
        UPDATE %s
        SET cleaner_id = $2
        WHERE request_id = $1
    `, svc.requestablename)

	_, err := svc.db.Exec(query, requestId, cleanerId)
	if err != nil {
		return err
	}
	return nil
}

func (svc postgresClient) GetRequestByClient(clientId string) (*[]domain.Request, error) {
	query := fmt.Sprintf(`
        SELECT request_id, client_id, cleaner_id, service_id, requested_date, status, created_at, updated_at
        FROM %s
        WHERE client_id = $1
    `, svc.requestablename)

	rows, err := svc.db.Query(query, clientId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []domain.Request
	for rows.Next() {
		var request domain.Request
		err := rows.Scan(
			&request.RequestId,
			&request.ClientId,
			&request.CleanerId,
			&request.ServiceId,
			&request.RequestedDate,
			&request.Status,
			&request.CreatedAt,
			&request.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &requests, nil
}

func (svc postgresClient) GetRequestByCleaner(cleanerId string) (*[]domain.Request, error) {
	query := fmt.Sprintf(`
        SELECT request_id, client_id, cleaner_id, service_id, requested_date, status, created_at, updated_at
        FROM %s
        WHERE cleaner_id = $1
    `, svc.requestablename)

	rows, err := svc.db.Query(query, cleanerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []domain.Request
	for rows.Next() {
		var request domain.Request
		err := rows.Scan(
			&request.RequestId,
			&request.ClientId,
			&request.CleanerId,
			&request.ServiceId,
			&request.RequestedDate,
			&request.Status,
			&request.CreatedAt,
			&request.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &requests, nil
}

func (svc postgresClient) CreateReview(review domain.Reviews) (*domain.Reviews, error) {
	query := fmt.Sprintf(`
        INSERT INTO %s (review_id, request_id, client_id, cleaner_id, rating, comment, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `, svc.reviewTablename)

	_, err := svc.db.Exec(query,
		review.ReviewId,
		review.RequestId,
		review.ClientId,
		review.CleanerId,
		review.Rating,
		review.Comment,
		review.CreatedAt,
		review.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return svc.GetReviewById(review.ReviewId)
}

func (svc postgresClient) GetReviewById(reviewId string) (*domain.Reviews, error) {
	query := fmt.Sprintf(`
        SELECT review_id, request_id, client_id, cleaner_id, rating, comment, created_at, updated_at
        FROM %s
        WHERE review_id = $1
    `, svc.reviewTablename)

	var review domain.Reviews
	err := svc.db.QueryRow(query, reviewId).Scan(
		&review.ReviewId,
		&review.RequestId,
		&review.ClientId,
		&review.CleanerId,
		&review.Rating,
		&review.Comment,
		&review.CreatedAt,
		&review.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (svc postgresClient) UpdateReview(review domain.Reviews) (*domain.Reviews, error) {
	query := fmt.Sprintf(`
        UPDATE %s
        SET request_id = $2, client_id = $3, cleaner_id = $4, rating = $5, comment = $6, updated_at = $7
        WHERE review_id = $1
    `, svc.reviewTablename)

	_, err := svc.db.Exec(query,
		review.ReviewId,
		review.RequestId,
		review.ClientId,
		review.CleanerId,
		review.Rating,
		review.Comment,
		review.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return svc.GetReviewById(review.ReviewId)
}

func (svc postgresClient) DeleteReview(reviewId string) error {
	query := fmt.Sprintf(`
        DELETE FROM %s
        WHERE review_id = $1
    `, svc.reviewTablename)

	_, err := svc.db.Exec(query, reviewId)
	if err != nil {
		return err
	}
	return nil
}

func (svc postgresClient) GetReviewByClient(clientId string) (*[]domain.Reviews, error) {
	query := fmt.Sprintf(`
        SELECT review_id, request_id, client_id, cleaner_id, rating, comment, created_at, updated_at
        FROM %s
        WHERE client_id = $1
    `, svc.reviewTablename)

	rows, err := svc.db.Query(query, clientId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []domain.Reviews
	for rows.Next() {
		var review domain.Reviews
		err := rows.Scan(
			&review.ReviewId,
			&review.RequestId,
			&review.ClientId,
			&review.CleanerId,
			&review.Rating,
			&review.Comment,
			&review.CreatedAt,
			&review.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &reviews, nil
}

func (svc postgresClient) GetReviewByCleaner(cleanerId string) (*[]domain.Reviews, error) {
	query := fmt.Sprintf(`
        SELECT review_id, request_id, client_id, cleaner_id, rating, comment, created_at, updated_at
        FROM %s
        WHERE cleaner_id = $1
    `, svc.reviewTablename)

	rows, err := svc.db.Query(query, cleanerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []domain.Reviews
	for rows.Next() {
		var review domain.Reviews
		err := rows.Scan(
			&review.ReviewId,
			&review.RequestId,
			&review.ClientId,
			&review.CleanerId,
			&review.Rating,
			&review.Comment,
			&review.CreatedAt,
			&review.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &reviews, nil
}
