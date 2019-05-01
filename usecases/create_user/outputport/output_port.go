package outputport

import "pumpkin/domain/model"

type CreateUserOutputPort interface {
	CreateUser(user *model.User) error
}
