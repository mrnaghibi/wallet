package service

import (
	"github.com/mrnaghibi/wallet/entity"
	"github.com/mrnaghibi/wallet/repository"
)

type WalletService struct {
	repo repository.WalletRepository
}

func WalletServiceProvider(repository repository.WalletRepository) WalletService {
	return WalletService{repo: repository}
}

func (s *WalletService) ChargeWallet(mobile string, amount float64) entity.Wallet {
	return s.repo.ChargeWallet(mobile, amount)
}
func (s *WalletService) Read(mobileNumber string) entity.Wallet {
	return s.repo.Read(mobileNumber)
}
