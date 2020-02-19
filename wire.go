//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mrnaghibi/wallet/controller"
	"github.com/mrnaghibi/wallet/repository"
	"github.com/mrnaghibi/wallet/service"
)

func initWalletController() controller.WalletController {
	wire.Build(repository.WalletRepositoryProvider,service.WalletServiceProvider,controller.WalletControllerProvider)
	return controller.WalletController{}
}
