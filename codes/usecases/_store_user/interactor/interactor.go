package interactor

import (
	"pumpkin/codes/domain/model"
	"pumpkin/codes/usecases/_store_user/inputport"
	"pumpkin/codes/usecases/_store_user/outputport"
)

type interactor struct {
	output outputport.StoreUserOutputPort
}

type StoreUserInteractor interface {
	inputport.StoreUserInputPort
}

// inherit input port and invoke presenter
func NewInteractor(presenter outputport.StoreUserOutputPort) StoreUserInteractor {
	return &interactor{presenter}
}

func (inter *interactor) StoreUser(user *model.User) error {
	return inter.output.StoreUser(user)
}
