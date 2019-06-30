package controllers

import (
	"pumpkin/codes/domain/model"
	"pumpkin/codes/usecases/db/_store_user/inputport"
)

type controller struct {
	inputPort inputport.StoreUserInputPort
}

type StoreUserController interface {
	StoreUser(*model.User) error
}

// Add gateways
// invoke Input port
// Use interactor as instance of InputPort
func NewController(inputPort inputport.StoreUserInputPort) StoreUserController {
	return &controller{inputPort}
}

func (userController *controller) StoreUser(pUser *model.User) error {
	// interactorに処理を委譲して対応するPresenterを呼び出してもらう
	return userController.inputPort.StoreUser(pUser)
}
