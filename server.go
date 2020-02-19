package main

import (
	router "github.com/mrnaghibi/wallet/http"
)

var httpRouter = router.NewMuxRouter()

func handleRequest() {

	walletController := initWalletController()

	httpRouter.POST("/api/wallets", walletController.ReadWallet)
	httpRouter.POST("/api/charge", walletController.ChargeWallet)
	httpRouter.SERVE(":2020")
}

func main() {
	handleRequest()
}
