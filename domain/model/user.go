package model

// Entity
import (
	"time"
)

type User struct {
	ID        uint       `json:"-" sql:"primary_key"`
	UserName  string     `json:"name"`
	Mail      string     `json:"email"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

type Users []User
