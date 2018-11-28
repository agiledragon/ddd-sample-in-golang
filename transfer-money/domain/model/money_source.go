package model

import (
	"github.com/agiledragon/ddd-dci-sample-in-golang/transfer-money/domain/model/base"
)

type moneySource struct {
	base.Role
	accountId string
	balance   *balance
	phone     *phone
}

func newMoneySource(accountId string, balance *balance, phone *phone) *moneySource {
	return &moneySource{accountId: accountId, balance: balance, phone: phone}
}

func (this *moneySource) TransferMoneyTo(dst *moneyDestination, amount uint) {
	if this.balance.get() < amount {
		panic("insufficient money!")
	}
	dst.transferMoneyFrom(this.accountId, amount)
	this.balance.decrease(amount)
	this.phone.sendTransferToMsg(dst.getAccountId(), amount)
}

func (this *moneySource) GetAmount() uint {
	return this.balance.get()
}
