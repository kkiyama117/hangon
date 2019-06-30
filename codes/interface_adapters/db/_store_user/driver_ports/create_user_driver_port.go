package driver_ports

import "pumpkin/codes/domain/model"

type DBOutput interface {
	StoreUser(*model.User) error
}
