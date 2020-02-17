package service

import (
	"../entity"
	"../repository"
)

type WalletService interface {
	CreateOrUpdate(*entity.Wallet) (*entity.Wallet, error)
	Read(string) (*entity.Wallet, error)
}

var (
	repo repository.Wallet
)

type service struct{}

func NewWalletService(repository repository.Wallet) WalletService {
	repo = repository
	return &service{}
}

func (*service) CreateOrUpdate(wallet *entity.Wallet) (*entity.Wallet, error) {
	return repo.InsertOrUpdate(wallet)
}
func (*service) Read(mobileNumber string) (*entity.Wallet, error) {
	return repo.Read(mobileNumber)
}
