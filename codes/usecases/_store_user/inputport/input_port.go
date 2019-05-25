package inputport

import "pumpkin/codes/domain/model"

type StoreUserInputPort interface {
	StoreUser(user *model.User) error
}