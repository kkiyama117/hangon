package inputport

import "pumpkin/codes/domain/model"

type GetUserInputPort interface {
	GetUser(user *model.User) error
}