package presenters

import (
	"encoding/json"

	"pumpkin/codes/domain/model"
	driver_ports2 "pumpkin/codes/interface_adapters/create_user/driver_ports"
	"pumpkin/codes/usecases/create_user/outputport"
)

type userPresenter struct {
	output driver_ports2.APIOutput
}

type CreateUserPresenter interface {
	// inherit OutputPort interface
	outputport.CreateUserOutputPort
}

func NewPresenter(output driver_ports2.APIOutput) CreateUserPresenter {
	return &userPresenter{output}
}

func (userPresenter *userPresenter) DoUsecase(pUser *model.User) error {
	// format user domain to []bytes for render
	users := model.Users{*pUser}
	res, err := json.Marshal(users)
	if err != nil {
		return err
	}
	// 描画処理
	err = userPresenter.output.ShowUser(res)
	if err != nil {
		return err
	}
	return nil
}
