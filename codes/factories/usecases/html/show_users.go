package html

import (
	"pumpkin/codes/interface_adapters/html/user/controllers"
	"pumpkin/codes/interface_adapters/html/user/driver_ports"
	"pumpkin/codes/interface_adapters/html/user/presenters"
	"pumpkin/codes/usecases/html/interactor"
)

func InjectedShowUsers(output driver_ports.APIOutput) controllers.ShowUsersController {
	pres := presenters.NewShowUsersPresenter(output)
	inter := interactor.NewShowUsersInteractor(pres)
	c := controllers.NewShowUsersController(inter)
	return c
}
