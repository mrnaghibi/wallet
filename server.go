package main

import (
	"os"
"log"
	"../controller"
	router "../http"
	"../repository"
	"../service"
)

var (
	walletRepository repository.Wallet           = repository.NewSliceRepository()
	walletService    service.WalletService       = service.NewWalletService(walletRepository)
	walletController controller.WalletController = controller.NewWalletController(walletService)
	httpRouter         router.Router                 = router.NewMuxRouter()
)

func handleRequest() {
	
	httpRouter.GET("/api/wallet", walletController.ReadWallet)
	httpRouter.POST("/api/discounts/consume",walletController.CreateOrUpdateWallet)
	log.Println(os.Getenv("PORT"))
	httpRouter.SERVE(os.Getenv("PORT"))
	//httpRouter.SERVE(":8000")
}

func main() {
	handleRequest()
}
