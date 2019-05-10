package presenters

import (
	"encoding/json"

	"pumpkin/domain/model"
	"pumpkin/external_interfaces"
	"pumpkin/usecases/create_user/outputport"
)

type userPresenter struct {
	output external_interfaces.Output
}

type CreateUserPresenter interface {
	// inherit OutputPort interface
	outputport.CreateUserOutputPort
}

func NewPresenter(framework external_interfaces.Output) CreateUserPresenter {
	return &userPresenter{framework}
}

func (userPresenter *userPresenter) DoUsecase(pUser *model.User) error {
	// format user domain to []bytes for render
	users:=model.Users{*pUser}
	res, err := json.Marshal(users)
	if err != nil {
		return err
	}
	// 描画処理
	// DI
	err = userPresenter.output.Push(res)
	if err != nil {
		return err
	}
	return nil
}
