package driver_ports

import "pumpkin/codes/domain/model"

type UserDBOutput interface {
	GetUser(*model.User) error
	GetUsers([]*model.User) error
	StoreUser(*model.User) error
}
