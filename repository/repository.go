package repository

import (
	"../entity"
)
type Wallet interface{
	InsertOrUpdate(*entity.Wallet) (*entity.Wallet,error)
	Read(string) (*entity.Wallet,error)
}