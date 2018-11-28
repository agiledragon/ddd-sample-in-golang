package model

import "github.com/agiledragon/ddd-dci-sample-in-golang/transfer-money/domain/model/base"

type Account struct {
	base.AggregateRoot
	balance          *balance
	phone            *phone
	MoneySource      *moneySource
	MoneyDestination *moneyDestination
}

func newAccount(accountId string, balance *balance, phone *phone) *Account {
	account := &Account{
		AggregateRoot: base.NewAggregateRoot(accountId),
		balance:       balance,
		phone:         phone,
	}
	account.MoneySource = newMoneySource(accountId, balance, phone)
	account.MoneyDestination = newMoneyDestination(accountId, balance, phone)
	return account
}
