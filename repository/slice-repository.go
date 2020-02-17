package repository

import (
	"errors"
	"math/rand"

	"../entity"
)

var (
	wallets  []entity.Wallet
)

type sliceRepository struct{}

func NewSliceRepository() Wallet {
	return &sliceRepository{}
}

func (*sliceRepository) InsertOrUpdate(wall *entity.Wallet) (*entity.Wallet, error) {

	for index, wallet := range wallets {
		if wallet.Mobile == mobileNumber {
			wallets = append(wallets[:index], wallets[index+1:]...)
			wallet.Balance += wall.Balance
			discounts = append(wallets, entity.Wallet{ID: wallet.ID, Mobile: wallet.Mobile, Balance: wallet.Balance})
			return &wallet , nil
		}
	}
	wall.ID = rand.Int63()
	wallets = append(wallets, *wall)
	return wall, nil
}
func (*sliceRepository) Read(mobileNumber string) (*entity.Wallet,error) {
	for _, wallet := range wallets {
		if wallet.Mobile == mobileNumber {
			return &wallet , nil
		}
	}
	return nil,errors.New("Wallet Doesn't Exist")
}
