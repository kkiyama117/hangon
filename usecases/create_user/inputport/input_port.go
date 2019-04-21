package inputport

import (
	"pumpkin/domain/model"
)

type CreateUserInputPort interface {
	CreateUser(u *model.User) error
}
