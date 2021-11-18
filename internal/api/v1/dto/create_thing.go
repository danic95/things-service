package dto

import "time"

// CreateThing represents the parameters for Swagger documentation
//
// swagger:parameters createThing
type CreateThing struct {
	// in:body
	Body struct {
		Name       string `json:"name" validate:"required"`
		Email      string `json:"email" validate:"email"`
		Phone      string `json:"phone" validate:"e164"`
		Birth_date string `json:"birth_date" validate:"required"`
		Start_day  int    `json:"start_day" validate:"required,numeric"`
	}
}

// CreateThingRS represents the response for Swagger documentation
//
// swagger:response CreateThingRS
type CreateThingRS struct {
	// in:body
	Body struct {
		ID         uint      `json:"id"`
		Name       string    `json:"name"`
		Email      string    `json:"email"`
		Phone      string    `json:"phone"`
		Birth_date time.Time `json:"birth_date"`
		Start_day  int       `json:"start_day"`
	}
}
