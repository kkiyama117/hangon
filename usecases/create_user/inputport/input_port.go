package inputport

import "pumpkin/domain/model"

type CreateUserInputPort interface {
	CreateUser(user *model.User) error
}
