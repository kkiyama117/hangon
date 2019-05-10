package presenters

import (
	"pumpkin/domain/model"
	"pumpkin/external_interfaces"
	"pumpkin/usecases/_store_user/outputport"
)

type userPresenter struct {
	output external_interfaces.Output
}

type StoreUserPresenter interface {
	// inherit OutputPort interface
	outputport.StoreUserOutputPort
}

func NewPresenter(output external_interfaces.Output) StoreUserPresenter {
	return &userPresenter{output}
}

func (userPresenter *userPresenter) StoreUser(pUser *model.User) error {
	user:=*pUser
	err := userPresenter.output.Push(user)
	if err != nil {
		return err
	}
	return nil
}