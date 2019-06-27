package outputport

import (
	"pumpkin/codes/domain/model"
)

type CreateUserOutputPort interface {
	DoUsecase(user *model.User) error
}
