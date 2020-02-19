package repository

import (
	"github.com/mrnaghibi/wallet/entity"
)

type WalletRepository struct{}

func WalletRepositoryProvider() WalletRepository {
	return WalletRepository{}
}

var userWallets = make(map[string]float64)

func (*WalletRepository) ChargeWallet(mobile string, amount float64) entity.Wallet {
	if _, ok := userWallets[mobile]; ok {
		userWallets[mobile] += amount
	} else {
		userWallets[mobile] = amount
	}

	return entity.Wallet{
		Mobile:  mobile,
		Balance: userWallets[mobile],
	}
}

func (*WalletRepository) Read(mobile string) entity.Wallet {
	return entity.Wallet{
		Mobile:  mobile,
		Balance: userWallets[mobile],
	}
}
