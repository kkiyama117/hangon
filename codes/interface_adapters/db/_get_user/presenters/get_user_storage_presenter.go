package presenters

import (
	"pumpkin/codes/domain/model"
	"pumpkin/codes/interface_adapters/db/_get_user/driver_ports"
	"pumpkin/codes/usecases/db/_get_user/outputport"
)

type userPresenter struct {
	output driver_ports.DBOutput
}

type GetUserPresenter interface {
	// inherit OutputPort interface
	outputport.GetUserOutputPort
}

func NewPresenter(output driver_ports.DBOutput) GetUserPresenter {
	return &userPresenter{output}
}

func (userPresenter *userPresenter) GetUser(pUser *model.User) error {
	err := userPresenter.output.GetUser(pUser)
	if err != nil {
		return err
	}
	return nil
}