package presenters

import (
	"encoding/json"

	"pumpkin/codes/domain/model"
	driver_ports2 "pumpkin/codes/interface_adapters/html/_show_user/driver_ports"
	"pumpkin/codes/usecases/html/_show_user/outputport"
)

type userPresenter struct {
	output driver_ports2.APIOutput
}

type ShowUserPresenter interface {
	// inherit OutputPort interface
	outputport.ShowUserOutputPort
}

func NewPresenter(output driver_ports2.APIOutput) ShowUserPresenter {
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
