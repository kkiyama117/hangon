package controllers

import (
	"pumpkin/domain/model"
	"pumpkin/usecases/create_user/inputport"
)

type controller struct {
	inputPort inputport.CreateUserInputPort
}
type CreateUserController interface {
	CreateUser(*model.User) error
}

// Add gateways
// invoke Input port
// Use interactor as instance of InputPort
func NewController(inputPort inputport.CreateUserInputPort) *controller {
	return &controller{inputPort}
}

func (controller *controller) CreateUser(pUser *model.User) error {
	// interactorに処理を委譲して対応するPresenterを呼び出してもらう
	return controller.inputPort.DoUsecase(pUser)
}

