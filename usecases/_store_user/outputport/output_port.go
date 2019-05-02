package outputport

import "pumpkin/domain/model"

type StoreUserOutputPort interface {
	StoreUser(user *model.User) error
}