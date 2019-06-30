package db

import (
	"pumpkin/codes/interface_adapters/db/user/controllers"
	"pumpkin/codes/interface_adapters/db/user/driver_ports"
	"pumpkin/codes/interface_adapters/db/user/presenters"
	"pumpkin/codes/usecases/db/interactor"
)

func InjectedStoreUser(output driver_ports.UserDBOutput) controllers.StoreUserController {
	pres := presenters.NewStoreUserPresenter(output)
	inter := interactor.NewStoreUserInteractor(pres)
	c := controllers.NewStoreUserController(inter)
	return c
}
