package interactor

import (
	"pumpkin/codes/domain/model"
	"pumpkin/codes/usecases/create_user/outputport"
)

type interactor struct {
	output outputport.CreateUserOutputPort
}

type CreateUserInteractor interface {
	inputport2.CreateUserInputPort
}

// inherit input port and invoke presenter
func NewInteractor(presenter outputport.CreateUserOutputPort) CreateUserInteractor {
	return &interactor{presenter}
}

func (inter *interactor) DoUsecase(user *model.User) error {
	return inter.output.DoUsecase(user)
}
