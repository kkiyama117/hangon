package usecases

import (
	controllers2 "pumpkin/codes/interface_adapters/create_user/controllers"
	driver_ports2 "pumpkin/codes/interface_adapters/create_user/driver_ports"
	presenters2 "pumpkin/codes/interface_adapters/create_user/presenters"
	"pumpkin/codes/usecases/create_user/interactor"
)

func InjectedCreateUser(output driver_ports2.APIOutput) controllers2.CreateUserController {
	pres := presenters2.NewPresenter(output)
	inter := interactor.NewInteractor(pres)
	c := controllers2.NewController(inter)
	return c
}
