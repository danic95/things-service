package service

import (
	"github.com/danic95/things-service/internal/things"
)

// Service is a struct able to access all data required
// to perform business logic functions
type Service struct {
	repo things.Repository
}

// New constructs and returns a Service struct
func New(repo things.Repository) things.Service {
	return &Service{
		repo: repo,
	}
}
