package db

import (
	"pumpkin/codes/interface_adapters/db/_get_user/controllers"
	"pumpkin/codes/interface_adapters/db/_get_user/driver_ports"
	"pumpkin/codes/interface_adapters/db/_get_user/presenters"
	"pumpkin/codes/usecases/db/_get_user/interactor"
)

func InjectedGetUser(output driver_ports.DBOutput) controllers.GetUserController {
	pres := presenters.NewPresenter(output)
	inter := interactor.NewInteractor(pres)
	c := controllers.NewController(inter)
	return c
}
