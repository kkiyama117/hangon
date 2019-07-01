package driver_ports

import "pumpkin/codes/domain/model"

type UserDBOutput interface {
	GetUser(*model.User) error
	GetUsers(model.Users) error
	StoreUser(*model.User) error
}
