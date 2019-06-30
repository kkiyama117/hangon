package interactor

import (
	"pumpkin/codes/domain/model"
	"pumpkin/codes/usecases/db/_get_user/inputport"
	"pumpkin/codes/usecases/db/_get_user/outputport"
)

type interactor struct {
	output outputport.GetUserOutputPort
}

type GetUserInteractor interface {
	inputport.GetUserInputPort
}

// inherit input port and invoke presenter
func NewInteractor(presenter outputport.GetUserOutputPort) GetUserInteractor {
	return &interactor{presenter}
}

func (inter *interactor) GetUser(user *model.User) error {
	return inter.output.GetUser(user)
}
