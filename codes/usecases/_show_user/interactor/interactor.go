package interactor

import (
	"pumpkin/codes/domain/model"
	"pumpkin/codes/usecases/_show_user/inputport"
	"pumpkin/codes/usecases/_show_user/outputport"
)

type interactor struct {
	output outputport.ShowUserOutputPort
}

type ShowUserInteractor interface {
	inputport.ShowUserInputPort
}

// inherit input port and invoke presenter
func NewInteractor(presenter outputport.ShowUserOutputPort) ShowUserInteractor {
	return &interactor{presenter}
}

func (inter *interactor) DoUsecase(user *model.User) error {
	return inter.output.DoUsecase(user)
}
