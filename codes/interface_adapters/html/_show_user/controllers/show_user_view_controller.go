package controllers

import (
	"pumpkin/codes/domain/model"
	"pumpkin/codes/usecases/html/_show_user/inputport"
)

type controller struct {
	inputPort inputport.ShowUserInputPort
}
type ShowUserController interface {
	ShowUser(*model.User) error
}

// Add gateways
// invoke Input port
// Use interactor as instance of InputPort
func NewController(inputPort inputport.ShowUserInputPort) ShowUserController {
	return &controller{inputPort}
}

func (controller *controller) ShowUser(pUser *model.User) error {
	// interactorに処理を委譲して対応するPresenterを呼び出してもらう
	return controller.inputPort.DoUsecase(pUser)
}
