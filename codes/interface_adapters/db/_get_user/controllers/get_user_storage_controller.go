package controllers

import (
	"pumpkin/codes/domain/model"
	"pumpkin/codes/usecases/db/_get_user/inputport"
)

type controller struct {
	inputPort inputport.GetUserInputPort
}

type GetUserController interface {
	GetUser(*model.User) error
}

// Add gateways
// invoke Input port
// Use interactor as instance of InputPort
func NewController(inputPort inputport.GetUserInputPort) GetUserController{
	return &controller{inputPort}
}

func (userController *controller) GetUser(pUser *model.User) error {
	// interactorに処理を委譲して対応するPresenterを呼び出してもらう
	return userController.inputPort.GetUser(pUser)
}
