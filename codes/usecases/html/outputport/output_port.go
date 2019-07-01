package outputport

import (
	"pumpkin/codes/domain/model"
)

type ShowUserOutputPort interface {
	DoUsecase(user *model.User) error
}

type ShowUsersOutputPort interface {
	DoUsecase(users model.Users) error
}
