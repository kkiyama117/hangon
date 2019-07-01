package controllers

import (
	"pumpkin/codes/domain/model"
	"pumpkin/codes/usecases/html/inputport"
)

type showUserController struct {
	inputPort inputport.ShowUserInputPort
}

type showUsersController struct {
	inputPort inputport.ShowUsersInputPort
}

type ShowUserController interface {
	ShowUser(*model.User) error
}

type ShowUsersController interface {
	ShowUsers(model.Users) error
}

// Add gateways
// invoke Input port
// Use interactor as instance of InputPort
func NewShowUserController(inputPort inputport.ShowUserInputPort) ShowUserController {
	return &showUserController{inputPort}
}

func (controller *showUserController) ShowUser(pUser *model.User) error {
	// interactorに処理を委譲して対応するPresenterを呼び出してもらう
	return controller.inputPort.DoUsecase(pUser)
}

func NewShowUsersController(inputPort inputport.ShowUsersInputPort) ShowUsersController {
	return &showUsersController{inputPort}
}
func (controller *showUsersController) ShowUsers(users model.Users) error {
	// interactorに処理を委譲して対応するPresenterを呼び出してもらう
	return controller.inputPort.DoUsecase(users)
}
