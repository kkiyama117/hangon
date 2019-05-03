package controllers

import (
	"pumpkin/domain/model"
	"pumpkin/usecases/create_user/inputport"
)

type userController struct {
	inputPort inputport.CreateUserInputPort
}

// Add gateways
// invoke Input port
// Use interactor as instance of InputPort
func NewController(inputPort inputport.CreateUserInputPort) inputport.CreateUserInputPort{
	return &userController{inputPort}
}

func (userController *userController) CreateUser(pUser *model.User) error {
	// interactorに処理を委譲して対応するPresenterを呼び出してもらう
	return userController.inputPort.CreateUser(pUser)
}

