package driver_ports

import "pumpkin/codes/domain/model"

type DBOutput interface {
	GetUser(*model.User) error
}
