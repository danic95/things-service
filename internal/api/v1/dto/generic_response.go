package dto

// Generic messages

// Generic success
// swagger:model genericSuccessRS
type GenericSuccessResp struct {
	// in: body
	Data interface{}
}

// Generic error
// swagger:model genericErrorRS
type GenericErrorResp struct {
	// Error message
	// in: body
	Body string `json:"error"`
}

// Bad request
// swagger:model badRequestRS
type BadRequestResp struct {
	// Error message
	// in: body
	Body string `json:"error"`
}

// Internal server error
// swagger:model serverErrorRS
type ServerErrorResp struct {
	// Error message
	// in: body
	Body string `json:"body"`
}
