package outputport

import "pumpkin/codes/domain/model"

type GetUserOutputPort interface {
	GetUser(user *model.User) error
}

type GetUsersOutputPort interface {
	GetUsers(users *model.Users) error
}

type StoreUserOutputPort interface {
	StoreUser(user *model.User) error
}
