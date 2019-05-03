package outputport

import (
	"pumpkin/domain/model"
)

type CreateUserOutputPort interface {
	DoUsecase(user *model.User) error
}
