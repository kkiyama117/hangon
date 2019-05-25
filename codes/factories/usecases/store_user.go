package usecases

import (
	controllers2 "pumpkin/codes/interface_adapters/_store_user/controllers"
	driver_ports2 "pumpkin/codes/interface_adapters/_store_user/driver_ports"
	presenters2 "pumpkin/codes/interface_adapters/_store_user/presenters"
	"pumpkin/codes/usecases/_store_user/interactor"
)

func InjectedStoreUser(output driver_ports2.DBOutput) controllers2.StoreUserController {
	pres := presenters2.NewPresenter(output)
	inter := interactor.NewInteractor(pres)
	c := controllers2.NewController(inter)
	return c
}
