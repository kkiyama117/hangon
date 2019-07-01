package controllers

import (
	"pumpkin/codes/domain/model"
	"pumpkin/codes/usecases/db/inputport"
)

type getUserController struct {
	inputPort inputport.GetUserInputPort
}

type getUsersController struct {
	inputPort inputport.GetUsersInputPort
}

type storeUserController struct {
	inputPort inputport.StoreUserInputPort
}

type GetUserController interface {
	GetUser(*model.User) error
}
type GetUsersController interface {
	GetUsers(*model.Users) error
}
type StoreUserController interface {
	StoreUser(*model.User) error
}

// Add gateways
// invoke Input port
// Use interactor as instance of InputPort
func NewGetUserController(inputPort inputport.GetUserInputPort) GetUserController {
	return &getUserController{inputPort}
}
func (userController *getUserController) GetUser(pUser *model.User) error {
	// interactorに処理を委譲して対応するPresenterを呼び出してもらう
	return userController.inputPort.GetUser(pUser)
}

func NewGetUsersController(inputPort inputport.GetUsersInputPort) GetUsersController {
	return &getUsersController{inputPort}
}
func (userController *getUsersController) GetUsers(pUsers *model.Users) error {
	return userController.inputPort.GetUsers(pUsers)
}

func NewStoreUserController(inputPort inputport.StoreUserInputPort) StoreUserController {
	return &storeUserController{inputPort}
}
func (userController *storeUserController) StoreUser(pUser *model.User) error {
	// interactorに処理を委譲して対応するPresenterを呼び出してもらう
	return userController.inputPort.StoreUser(pUser)
}
