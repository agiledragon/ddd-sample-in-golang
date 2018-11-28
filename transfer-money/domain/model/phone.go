package model

import (
	"fmt"
	"github.com/agiledragon/ddd-dci-sample-in-golang/transfer-money/domain/model/base"
)

type phone struct {
	base.ValueObject
	accountId   string
	phoneNumber string
}

func newPhone(accountId, phoneNumber string) *phone {
	return &phone{accountId: accountId, phoneNumber: phoneNumber}
}

func (this *phone) sendTransferToMsg(dstAccountId string, amount uint) {
	msg := fmt.Sprintf("%s send %d money to %s", this.accountId, amount, dstAccountId)
	GetPhoneProvider().Send(this.phoneNumber, msg)
}

func (this *phone) sendTransferFromMsg(srcAccountId string, amount uint) {
	msg := fmt.Sprintf("%s receive %d money from %s", this.accountId, amount, srcAccountId)
	GetPhoneProvider().Send(this.phoneNumber, msg)
}
