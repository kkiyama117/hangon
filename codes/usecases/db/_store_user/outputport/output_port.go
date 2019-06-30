package outputport

import "pumpkin/codes/domain/model"

type StoreUserOutputPort interface {
	StoreUser(user *model.User) error
}