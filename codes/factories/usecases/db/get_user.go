package db

import (
	"pumpkin/codes/interface_adapters/db/user/controllers"
	"pumpkin/codes/interface_adapters/db/user/driver_ports"
	"pumpkin/codes/interface_adapters/db/user/presenters"
	"pumpkin/codes/usecases/db/interactor"
)

func InjectedGetUser(output driver_ports.UserDBOutput) controllers.GetUserController {
	pres := presenters.NewGetUserPresenter(output)
	inter := interactor.NewGetUserInteractor(pres)
	c := controllers.NewGetUserController(inter)
	return c
}
