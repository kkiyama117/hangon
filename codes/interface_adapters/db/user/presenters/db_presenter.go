package presenters

import (
	"pumpkin/codes/domain/model"
	"pumpkin/codes/interface_adapters/db/user/driver_ports"
	"pumpkin/codes/usecases/db/outputport"
)

/* get user */
type getUserPresenter struct {
	output driver_ports.UserDBOutput
}

type GetUserPresenter interface {
	// inherit OutputPort interface
	outputport.GetUserOutputPort
}

func NewGetUserPresenter(output driver_ports.UserDBOutput) GetUserPresenter {
	return &getUserPresenter{output}
}

func (userPresenter *getUserPresenter) GetUser(pUser *model.User) error {
	err := userPresenter.output.GetUser(pUser)
	if err != nil {
		return err
	}
	return nil
}

/* get users */
type getUsersPresenter struct {
	output driver_ports.UserDBOutput
}

type GetUsersPresenter interface {
	// inherit OutputPort interface
	outputport.GetUsersOutputPort
}

func NewGetUsersPresenter(output driver_ports.UserDBOutput) GetUsersPresenter {
	return &getUsersPresenter{output}
}

func (userPresenter *getUsersPresenter) GetUsers(users model.Users) error {
	err := userPresenter.output.GetUsers(users)
	if err != nil {
		return err
	}
	return nil
}

/* store user */
type storeUserPresenter struct {
	output driver_ports.UserDBOutput
}

type StoreUserPresenter interface {
	// inherit OutputPort interface
	outputport.StoreUserOutputPort
}

func NewStoreUserPresenter(output driver_ports.UserDBOutput) StoreUserPresenter {
	return &storeUserPresenter{output}
}

func (userPresenter *storeUserPresenter) StoreUser(pUser *model.User) error {
	err := userPresenter.output.StoreUser(pUser)
	if err != nil {
		return err
	}
	return nil
}
