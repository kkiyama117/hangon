package inputport

import "pumpkin/domain/model"

type StoreUserInputPort interface {
	StoreUser(user *model.User) error
}