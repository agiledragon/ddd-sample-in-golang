package model

type AccountFactory struct {
}

func (this AccountFactory) Create(accountId string, amount uint, phoneNumber string) *Account {
	b := newBalance(amount)
	p := newPhone(accountId, phoneNumber)
	return newAccount(accountId, b, p)
}
