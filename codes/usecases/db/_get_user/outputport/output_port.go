package outputport

import "pumpkin/codes/domain/model"

type GetUserOutputPort interface {
	GetUser(user *model.User) error
}