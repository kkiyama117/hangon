package html

import (
	"pumpkin/codes/interface_adapters/html/user/controllers"
	"pumpkin/codes/interface_adapters/html/user/driver_ports"
	"pumpkin/codes/interface_adapters/html/user/presenters"
	"pumpkin/codes/usecases/html/interactor"
)

func InjectedShowUser(output driver_ports.APIOutput) controllers.ShowUserController {
	pres := presenters.NewShowUserPresenter(output)
	inter := interactor.NewShowUserInteractor(pres)
	c := controllers.NewShowUserController(inter)
	return c
}
