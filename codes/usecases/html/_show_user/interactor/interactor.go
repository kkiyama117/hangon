package interactor

import (
	"pumpkin/codes/domain/model"
	inputport2 "pumpkin/codes/usecases/html/_show_user/inputport"
	outputport2 "pumpkin/codes/usecases/html/_show_user/outputport"
)

type interactor struct {
	output outputport2.ShowUserOutputPort
}

type ShowUserInteractor interface {
	inputport2.ShowUserInputPort
}

// inherit input port and invoke presenter
func NewInteractor(presenter outputport2.ShowUserOutputPort) ShowUserInteractor {
	return &interactor{presenter}
}

func (inter *interactor) DoUsecase(user *model.User) error {
	return inter.output.DoUsecase(user)
}
