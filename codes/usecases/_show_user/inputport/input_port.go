package inputport

import (
	"pumpkin/codes/domain/model"
)

type ShowUserInputPort interface {
	DoUsecase(user *model.User) error
}
