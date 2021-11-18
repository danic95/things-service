package entity

import "time"

// Thing is db record of thing
type Thing struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Birth_date time.Time `json:"birth_date"`
	Start_day  int       `json:"start_day"`
}
