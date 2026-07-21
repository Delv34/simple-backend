package models

type State struct {
	ID    int32  `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Color string `json:"color" db:"color"`
}