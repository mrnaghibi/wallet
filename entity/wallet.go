package entity

type Wallet struct{
	ID int64 `json:id`
	Mobile string `json:"mobile"`
	Balance int64 `json:"balance"`
}