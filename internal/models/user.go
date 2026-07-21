package models

import "time"

// Строки расположены в таком порядке, для того чтобы оптимизировать приложение.
// Этот метод называется struct packing

type User struct {
	CreatedAt time.Time		`json:"created_at" db:"created_at"` //теги для json и db
	Username string			`json:"username" db:"username"` 
	Password string			`json:"password" db:"password"`
	ID uint32				`json:"id" db:"id"`
}