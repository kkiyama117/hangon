package db

import (
	"pumpkin/codes/interface_adapters/db/user/controllers"
	"pumpkin/codes/interface_adapters/db/user/driver_ports"
	"pumpkin/codes/interface_adapters/db/user/presenters"
	"pumpkin/codes/usecases/db/interactor"
)

func InjectedGetUsers(output driver_ports.UserDBOutput) controllers.GetUsersController {
	pres := presenters.NewGetUsersPresenter(output)
	inter := interactor.NewGetUsersInteractor(pres)
	c := controllers.NewGetUsersController(inter)
	return c
}
