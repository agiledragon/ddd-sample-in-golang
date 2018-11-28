package service

import (
	"github.com/agiledragon/ddd-dci-sample-in-golang/transfer-money/domain/service"
)

func CreateAccount(accountId string, amount uint, phoneNumber string) {
	service.GetAccountService().Create(accountId, amount, phoneNumber)
}

func TransferMoney(srcAccountId, dstAccountId string, amount uint) {
	service.GetAccountService().TransferMoney(srcAccountId, dstAccountId, amount)
}

func GetAmount(accountId string) (uint, error) {
	return service.GetAccountService().GetAmount(accountId)
}

func DestroyAccount(accountId string) {
	service.GetAccountService().Destroy(accountId)
}
