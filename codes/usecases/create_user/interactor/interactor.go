package interactor

import (
	"pumpkin/codes/domain/model"
	inputport2 "pumpkin/codes/usecases/create_user/inputport"
	outputport2 "pumpkin/codes/usecases/create_user/outputport"
)

type interactor struct {
	output outputport2.CreateUserOutputPort
}

type CreateUserInteractor interface {
	inputport2.CreateUserInputPort
}

// inherit input port and invoke presenter
func NewInteractor(presenter outputport2.CreateUserOutputPort) CreateUserInteractor {
	return &interactor{presenter}
}

func (inter *interactor) DoUsecase(user *model.User) error {
	return inter.output.DoUsecase(user)
}
