package models

import (
	"time"
)

type User struct {
	ID        int        `json:"id" db:"id"`
	Email     string     `json:"email" db:"email"`
	Name      string     `json:"name" db:"name"`
	CreatedAt *time.Time `json:"createdAt" db:"createdat"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updatedat"`
}
