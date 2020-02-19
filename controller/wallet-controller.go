package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mrnaghibi/wallet/errors"
	"github.com/mrnaghibi/wallet/service"
)

const (
	discountAmount float64 = 1000000
)

type mobileRequestModel struct {
	Mobile string `json:"mobile"`
}

type WalletController struct {
	service service.WalletService
}

func WalletControllerProvider(walletService service.WalletService) WalletController {
	return WalletController{service: walletService}
}

func (c *WalletController) ChargeWallet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var requestModel mobileRequestModel
	err := json.NewDecoder(request.Body).Decode(&requestModel)
	if err != nil {
		log.Printf("Wallet %v not charged: %v", requestModel.Mobile, err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiecError{Message: err.Error()})
		return
	}
	_ = c.service.ChargeWallet(requestModel.Mobile, discountAmount)
	response.WriteHeader(http.StatusNoContent)
}

func (c *WalletController) ReadWallet(response http.ResponseWriter, request *http.Request) {
	var requestModel mobileRequestModel
	err := json.NewDecoder(request.Body).Decode(&requestModel)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(errors.ServiecError{Message: "Bad Request!"})
		return
	}
	wallet := c.service.Read(requestModel.Mobile)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(wallet.Balance)
}
