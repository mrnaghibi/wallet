package repository

import (
	"github.com/mrnaghibi/wallet/entity"
)
type Wallet interface{
	InsertOrUpdate(*entity.Wallet) (*entity.Wallet,error)
	Read(string) (*entity.Wallet,error)
}