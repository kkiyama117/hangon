package model

// Entity
import (
	"time"
)

type User struct {
	Id        uint       `json:"-"`
	UserName  string     `json:"name"`
	Mail      string     `json:"email"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type Users []User
