package presenters

import (
	"pumpkin/codes/domain/model"
	driver_ports2 "pumpkin/codes/interface_adapters/db/_store_user/driver_ports"
	"pumpkin/codes/usecases/db/_store_user/outputport"
)

type userPresenter struct {
	output driver_ports2.DBOutput
}

type StoreUserPresenter interface {
	// inherit OutputPort interface
	outputport.StoreUserOutputPort
}

func NewPresenter(output driver_ports2.DBOutput) StoreUserPresenter {
	return &userPresenter{output}
}

func (userPresenter *userPresenter) StoreUser(pUser *model.User) error {
	err := userPresenter.output.StoreUser(pUser)
	if err != nil {
		return err
	}
	return nil
}