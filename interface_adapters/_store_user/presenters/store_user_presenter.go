package presenters

import (
	"pumpkin/domain/model"
	output "pumpkin/external_interfaces/common"
	"pumpkin/usecases/_store_user/outputport"
)

type userPresenter struct {
	output output.Output
}

type StoreUserPresenter interface {
	// inherit OutputPort interface
	outputport.StoreUserOutputPort
}

func NewPresenter(output output.Output) StoreUserPresenter {
	return &userPresenter{output}
}

func (userPresenter *userPresenter) StoreUser(pUser *model.User) error {
	user:=*pUser
	// 描画処理
	err := userPresenter.output.Push(user)
	if err != nil {
		return err
	}
	return nil
}