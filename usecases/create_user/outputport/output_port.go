package outputport

import "pumpkin/domain/model"

type CreateUserOutputPort interface {
	DoMethod(user *model.User) error
}
