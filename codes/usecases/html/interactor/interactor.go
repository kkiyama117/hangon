package interactor

import (
	"pumpkin/codes/domain/model"
	"pumpkin/codes/usecases/html/inputport"
	"pumpkin/codes/usecases/html/outputport"
)

/* Show User */
type showUserInteractor struct {
	output outputport.ShowUserOutputPort
}

type ShowUserInteractor interface {
	inputport.ShowUserInputPort
}

// inherit input port and invoke presenter
func NewShowUserInteractor(presenter outputport.ShowUserOutputPort) ShowUserInteractor {
	return &showUserInteractor{presenter}
}

func (inter *showUserInteractor) DoUsecase(user *model.User) error {
	return inter.output.DoUsecase(user)
}

/* Show Users */
type showUsersInteractor struct {
	output outputport.ShowUsersOutputPort
}

type ShowUsersInteractor interface {
	inputport.ShowUsersInputPort
}

// inherit input port and invoke presenter
func NewShowUsersInteractor(presenter outputport.ShowUsersOutputPort) ShowUsersInteractor {
	return &showUsersInteractor{presenter}
}

func (inter *showUsersInteractor) DoUsecase(users model.Users) error {
	return inter.output.DoUsecase(users)
}
