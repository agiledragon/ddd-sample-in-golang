package model

import (
	"github.com/agiledragon/ddd-dci-sample-in-golang/transfer-money/domain/model/base"
)

type moneyDestination struct {
	base.Role
	accountId string
	balance   *balance
	phone     *phone
}

func newMoneyDestination(accountId string, balance *balance, phone *phone) *moneyDestination {
	return &moneyDestination{accountId: accountId, balance: balance, phone: phone}
}

func (this *moneyDestination) getAccountId() string {
	return this.accountId
}

func (this *moneyDestination) transferMoneyFrom(srcAccountId string, amount uint) {
	this.balance.increase(amount)
	this.phone.sendTransferFromMsg(srcAccountId, amount)
}
