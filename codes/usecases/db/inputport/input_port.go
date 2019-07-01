package inputport

import "pumpkin/codes/domain/model"

type GetUserInputPort interface {
	GetUser(user *model.User) error
}

type GetUsersInputPort interface {
	GetUsers(pUsers *model.Users) error
}

type StoreUserInputPort interface {
	StoreUser(user *model.User) error
}
