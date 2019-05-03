package inputport

import (
	"pumpkin/domain/model"
)

type CreateUserInputPort interface {
	DoMethod(user *model.User) error
}
