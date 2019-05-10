package usecases

import (
	"pumpkin/external_interfaces"
	"pumpkin/interface_adapters/_store_user/controllers"
	"pumpkin/interface_adapters/_store_user/presenters"
	"pumpkin/usecases/_store_user/interactor"
)


func InjectedStoreUser(output external_interfaces.Output) controllers.StoreUserController{
	pres := presenters.NewPresenter(output)
	inter := interactor.NewInteractor(pres)
	c := controllers.NewController(inter)
	return c
}
