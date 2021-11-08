package v1

import (
	validator2 "github.com/sanservices/apicore/validator"

	"github.com/danic95/things-service/internal/things"
	"github.com/danic95/things-service/goutils/settings"
)

// Handler handles v1 routes
type Handler struct {
	cfg       *settings.Settings
	service   things.Service
	validator *validator2.Validator
}

// NewHandler initialize main *Handler
func NewHandler(cfg *settings.Settings, svc things.Service, validator *validator2.Validator) *Handler {

	return &Handler{
		cfg:       cfg,
		service:   svc,
		validator: validator,
	}
}
