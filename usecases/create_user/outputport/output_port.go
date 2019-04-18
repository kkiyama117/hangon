package outputport

import "pumpkin/domain/model"

type CreateUserOutputPort interface {
	CreateUser(us *model.User) *model.Status
}
