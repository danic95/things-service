package entity

// Things is a collection of db records of things
type Things struct {
	Body []Thing `json:"thing"`
}
