package driver_ports

import "pumpkin/domain/model"

type DBOutput interface {
	StoreUser(*model.User) error
}
