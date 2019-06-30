package interactor

import (
	"pumpkin/codes/domain/model"
	inputport2 "pumpkin/codes/usecases/db/_store_user/inputport"
	outputport2 "pumpkin/codes/usecases/db/_store_user/outputport"
)

type interactor struct {
	output outputport2.StoreUserOutputPort
}

type StoreUserInteractor interface {
	inputport2.StoreUserInputPort
}

// inherit input port and invoke presenter
func NewInteractor(presenter outputport2.StoreUserOutputPort) StoreUserInteractor {
	return &interactor{presenter}
}

func (inter *interactor) StoreUser(user *model.User) error {
	return inter.output.StoreUser(user)
}
