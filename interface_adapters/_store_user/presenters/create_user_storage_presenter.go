package presenters

import (
	"pumpkin/domain/model"
	"pumpkin/interface_adapters/_store_user/driver_ports"
	"pumpkin/usecases/_store_user/outputport"
)

type userPresenter struct {
	output driver_ports.DBOutput
}

type StoreUserPresenter interface {
	// inherit OutputPort interface
	outputport.StoreUserOutputPort
}

func NewPresenter(output driver_ports.DBOutput) StoreUserPresenter {
	return &userPresenter{output}
}

func (userPresenter *userPresenter) StoreUser(pUser *model.User) error {
	err := userPresenter.output.StoreUser(pUser)
	if err != nil {
		return err
	}
	return nil
}