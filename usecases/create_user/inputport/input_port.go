package inputport

import (
	"pumpkin/domain/model"
)

type CreateUserInputPort interface {
	DoUsecase(user *model.User) error
}
