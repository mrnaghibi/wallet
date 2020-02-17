package controller

import (
	"net/http"

	"encoding/json"

	"../entity"
	"../errors"
	"../service"
)

const(
	discountAmount = 1000000
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
	Mobile   string `json:"mobile"`
}

var (
	walletService service.walletService
)

type controller struct{}

func NewWalletController(walletSRV service.walletService) WalletController {
	walletService = walletSRV
	return &controller{}
}

func (*controller) CreateOrUpdateWallet(response http.ResponseWriter, request *http.Request) {
	
	response.Header().Set("Content-Type","application/json")

	var body bodyDiscount
	json.NewDecoder(request.Body).Decode(&body)

	jsonValue,err  := ioutil.ReadAll(request.Body)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiecError{Message: err.Error()})
		return
	}
	res , _ := http.Post("http://localhost:8000/api/discounts/consume","application/json",bytes.NewBuffer(jsonValue))
	if res.StatusCode == 204{
		var wallet entity.Wallet{
			Mobile: body.Mobile
			Balance : discountAmount
		}
		walletService.CreateOrUpdate(&wallet)
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(wallet)
	}else{
		http.Error(response, "{}", http.StatusBadRequest)
		return
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
	err1 := walletService.Read(body.Mobile)
	if err1 != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(errors.ServiecError{Message: err1.Error()})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(wallet)
}
