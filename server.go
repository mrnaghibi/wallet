package main

import (
	router "github.com/mrnaghibi/wallet/http"
	"os"
)

var httpRouter = router.NewMuxRouter()

func handleRequest() {

	walletController := initWalletController()

	httpRouter.POST("/api/wallets", walletController.ReadWallet)
	httpRouter.POST("/api/charge", walletController.ChargeWallet)
	httpRouter.SERVE(os.Getenv("PORT"))
}

func main() {
	handleRequest()
}
