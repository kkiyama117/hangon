package interactor

import (
	"pumpkin/codes/domain/model"
	"pumpkin/codes/usecases/db/inputport"
	"pumpkin/codes/usecases/db/outputport"
)

/* Get User */
type getUserInteractor struct {
	output outputport.GetUserOutputPort
}

type GetUserInteractor interface {
	inputport.GetUserInputPort
}

// inherit input port and invoke presenter
func NewGetUserInteractor(presenter outputport.GetUserOutputPort) GetUserInteractor {
	return &getUserInteractor{presenter}
}

func (inter *getUserInteractor) GetUser(user *model.User) error {
	return inter.output.GetUser(user)
}

/* Get Users */
type getUsersInteractor struct {
	output outputport.GetUsersOutputPort
}

type GetUsersInteractor interface {
	inputport.GetUsersInputPort
}

// inherit input port and invoke presenter
func NewGetUsersInteractor(presenter outputport.GetUsersOutputPort) GetUsersInteractor {
	return &getUsersInteractor{presenter}
}

func (inter *getUsersInteractor) GetUsers(users model.Users) error {
	return inter.output.GetUsers(users)
}

/* Store Users */
type storeUserInteractor struct {
	output outputport.StoreUserOutputPort
}

type StoreUserInteractor interface {
	inputport.StoreUserInputPort
}

// inherit input port and invoke presenter
func NewStoreUserInteractor(presenter outputport.StoreUserOutputPort) StoreUserInteractor {
	return &storeUserInteractor{presenter}
}

func (inter *storeUserInteractor) StoreUser(user *model.User) error {
	return inter.output.StoreUser(user)
}
