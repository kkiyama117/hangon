package usecases

import (
	"pumpkin/codes/interface_adapters/_store_user/controllers"
	"pumpkin/codes/interface_adapters/_store_user/driver_ports"
	"pumpkin/codes/interface_adapters/_store_user/presenters"
	"pumpkin/codes/usecases/_store_user/interactor"
)

func InjectedStoreUser(output driver_ports.DBOutput) controllers.StoreUserController {
	pres := presenters.NewPresenter(output)
	inter := interactor.NewInteractor(pres)
	c := controllers.NewController(inter)
	return c
}
