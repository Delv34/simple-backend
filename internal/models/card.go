package models

import "time"

type Card struct {
	CreatedAt time.Time		`json:"created_at" db:"created_at"`
	Title string			`json:"title" db:"title"`
	Description *string		`json:"description,omitempty" db:"description"`
	DoBefore *time.Time 	`json:"do_before,omitempty" db:"do_before"`
	ID uint32				`json:"id" db:"id"`
	UserID uint32			`json:"user_id" db:"user_id"`
	StateID uint32			`json:"state_id" db:"state_id"`
}