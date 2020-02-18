package repository

import (
	"errors"
	"github.com/mrnaghibi/wallet/entity"
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
		if wallet.Mobile == wall.Mobile {
			wallets[index].Balance += wall.Balance
			return &wallets[index] , nil
		}
	}
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
