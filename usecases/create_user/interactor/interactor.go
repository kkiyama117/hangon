package interactor

import (
	"pumpkin/domain/model"
	"pumpkin/usecases/create_user/inputport"
	"pumpkin/usecases/create_user/outputport"
)

type interactor struct {
	inputport.CreateUserInputPort
	output          outputport.CreateUserOutputPort
}

type Interactor interface {
	inputport.CreateUserInputPort
}

func NewInteractor(controller inputport.CreateUserInputPort, presenter outputport.CreateUserOutputPort) Interactor{
	return &interactor{controller, presenter}
}

func (inter *interactor) CreateUser(user *model.User) *model.Status {
	return inter.output.CreateUser(user)
}
