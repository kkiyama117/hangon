package inputport

import (
	"pumpkin/codes/domain/model"
)

type CreateUserInputPort interface {
	DoUsecase(user *model.User) error
}
