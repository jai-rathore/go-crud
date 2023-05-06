package models

import (
	"time"
)

type List struct {
	ID          int        `db:"id" json:"id"`
	Title       string     `db:"title" json:"title"`
	Description string     `db:"description" json:"description"`
	IsPublic    bool       `db:"ispublic" json:"isPublic"`
	UserID      int        `db:"userid" json:"userId"`
	CreatedAt   *time.Time `db:"createdat" json:"createdAt"`
	UpdatedAt   *time.Time `db:"updatedat" json:"updatedAt"`
}
