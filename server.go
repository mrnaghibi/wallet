package main

import (
	"os"
	"github.com/mrnaghibi/wallet/controller"
	router "github.com/mrnaghibi/wallet/http"
	"github.com/mrnaghibi/wallet/repository"
	"github.com/mrnaghibi/wallet/service"
)

var (
	walletRepository repository.Wallet           = repository.NewSliceRepository()
	walletService    service.WalletService       = service.NewWalletService(walletRepository)
	walletController controller.WalletController = controller.NewWalletController(walletService)
	httpRouter       router.Router               = router.NewMuxRouter()
)

func handleRequest() {
	httpRouter.POST("/api/wallets", walletController.ReadWallet)
	httpRouter.POST("/api/discounts", walletController.CreateOrUpdateWallet)
	httpRouter.SERVE(os.Getenv("PORT"))
}

func main() {
	handleRequest()
}
