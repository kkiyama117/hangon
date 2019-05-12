package model

// Entity
import (
	"time"
)

type User struct {
	ID        uint       `json:"id" sql:"primary_key"`
	UserName  string     `json:"name"`
	Mail      string     `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type Users []User
