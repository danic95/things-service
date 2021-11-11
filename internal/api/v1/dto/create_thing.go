package dto

// CreateThing represents the parameters for Swagger documentation
//
// swagger:parameters createThing
type CreateThing struct {
	// in:body
	Body struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
}

// CreateThingRS represents the response for Swagger documentation
//
// swagger:response CreateThingRS
type CreateThingRS struct {
	// in:body
	Body struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
}
