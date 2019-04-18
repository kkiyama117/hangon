package repository

import "pumpkin/domain/model"

type UserRepository interface {
	Store(user *model.User) error
}
