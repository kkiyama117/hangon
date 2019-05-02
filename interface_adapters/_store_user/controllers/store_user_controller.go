package controllers

import (
	"pumpkin/domain/model"
	"pumpkin/usecases/_store_user/inputport"
)

type userController struct {
	inputPort inputport.StoreUserInputPort
}

// Add gateways
// invoke Input port
// Use interactor as instance of InputPort
func NewController(inputPort inputport.StoreUserInputPort) *userController {
	return &userController{inputPort}
}

func (userController *userController) StoreUser(pUser *model.User) error {
	// interactorに処理を委譲して対応するPresenterを呼び出してもらう
	return userController.inputPort.StoreUser(pUser)
}
