package dto

import "github.com/danic95/things-service/internal/things/entity"

// disclaimer: there are non-ok json tag values for some of the *RQ structs, because they are in path or somewhere else, so it is not actual json

// swagger:model getThingsRS
type getThingsRS struct {
	// in: body
	Thing *entity.Things `json:"thing"`
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
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	/*Body struct {
		Thing *entity.Thing `json:"thing"`
	}*/
}
