package dto

import (
	"time"

	"github.com/danic95/things-service/internal/things/entity"
)

// disclaimer: there are non-ok json tag values for some of the *RQ structs, because they are in path or somewhere else, so it is not actual json

// swagger:model getThingsRS
type getThingsRS struct {
	// in: body
	Body struct {
		Things []entity.Thing `json:"things"`
	} `json:"data"`
}

// swagger:parameters getThingRQ
type getThingRQ struct {
	// in: path
	ID uint `json:"id"`
}

// swagger:model getThingRS
type getThingRS struct {
	// in: body
	Body struct {
		ID         uint      `json:"id"`
		Name       string    `json:"name"`
		Email      string    `json:"email"`
		Phone      string    `json:"phone"`
		Birth_date time.Time `json:"birth_date"`
		Start_day  int       `json:"start_day"`
	} `json:"data"`

	/*Body struct {
		Thing *entity.Thing `json:"thing"`
	}*/
}
