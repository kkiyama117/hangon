package presenters

import (
	"encoding/json"

	"pumpkin/codes/domain/model"
	"pumpkin/codes/interface_adapters/html/user/driver_ports"
	"pumpkin/codes/usecases/html/outputport"
)

/* show user */
type showUserPresenter struct {
	output driver_ports.APIOutput
}

type ShowUserPresenter interface {
	// inherit OutputPort interface
	outputport.ShowUserOutputPort
}

func NewShowUserPresenter(output driver_ports.APIOutput) ShowUserPresenter {
	return &showUserPresenter{output}
}

func (userPresenter *showUserPresenter) DoUsecase(pUser *model.User) error {
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

/* show users */
type showUsersPresenter struct {
	output driver_ports.APIOutput
}

type ShowUsersPresenter interface {
	// inherit OutputPort interface
	outputport.ShowUsersOutputPort
}

func NewShowUsersPresenter(output driver_ports.APIOutput) ShowUsersPresenter {
	return &showUsersPresenter{output}
}

func (userPresenter *showUsersPresenter) DoUsecase(users model.Users) error {
	// format user domain to []bytes for render
	res, err := json.Marshal(users)
	if err != nil {
		return err
	}
	// 描画処理
	err = userPresenter.output.ShowUsers(res)
	if err != nil {
		return err
	}
	return nil
}