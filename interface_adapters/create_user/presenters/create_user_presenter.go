package presenters

import (
	"encoding/json"
	"pumpkin/domain/model"
	output "pumpkin/external_interfaces/common"
	"pumpkin/usecases/create_user/outputport"
)

type userPresenter struct {
	output output.Output
}

type CreateUserPresenter interface {
	// inherit OutputPort interface
	outputport.CreateUserOutputPort
}

func NewPresenter(output output.Output) CreateUserPresenter {
	return &userPresenter{output}
}

func (userPresenter *userPresenter) CreateUser(us *model.User) error {
	// format user domain to []bytes for render
	jsonBytes, err := json.Marshal(us)
	if err != nil {
		return err
	}
	// 描画処理
	err = userPresenter.output.Push(jsonBytes)
	if err != nil {
		return err
	}
	return nil
}
