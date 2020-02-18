package controller

import (
	"os"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mrnaghibi/wallet/entity"
	"github.com/mrnaghibi/wallet/errors"
	"github.com/mrnaghibi/wallet/service"
)

const (
	discountAmount        = 1000000
)

type WalletController interface {
	CreateOrUpdateWallet(response http.ResponseWriter, request *http.Request)
	ReadWallet(response http.ResponseWriter, request *http.Request)
}
type bodyDiscount struct {
	Mobile   string `json:"mobile"`
	Discount string `json:"discount"`
}

type bodyMobile struct {
	Mobile string `json:"mobile"`
}

var (
	walletService service.WalletService
)

type controller struct{}

func NewWalletController(walletSRV service.WalletService) WalletController {
	walletService = walletSRV
	return &controller{}
}

func (*controller) CreateOrUpdateWallet(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")
	jsonValue, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiecError{Message: err.Error()})
		return
	}
	res, err3 := http.Post(os.Getenv("BASEURL")+"/api/discounts/consume", "application/json", bytes.NewBuffer(jsonValue))
	if err3 != nil {
		http.Error(response, err3.Error(), http.StatusInternalServerError)
		return
	}
	if res.StatusCode == 200 {
		var body bodyDiscount
		err1 := json.Unmarshal(jsonValue, &body)
		if err1 != nil {
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode(errors.ServiecError{Message: err1.Error()})
			return
		}
		wallet := entity.Wallet{
			Mobile:  body.Mobile,
			Balance: discountAmount,
		}
		result, err2 := walletService.CreateOrUpdate(&wallet)
		if err2 != nil {
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode(errors.ServiecError{Message: err1.Error()})
			return
		}
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(result)
	} else {
		response.WriteHeader(http.StatusForbidden)
		json.NewEncoder(response).Encode(errors.ServiecError{Message: "Error"})
	}
}
func (*controller) ReadWallet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var body bodyMobile
	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiecError{Message: "Unmarshalling Data"})
		return
	}

	wallet, err1 := walletService.Read(body.Mobile)
	if err1 != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(errors.ServiecError{Message: err1.Error()})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(wallet)
}
