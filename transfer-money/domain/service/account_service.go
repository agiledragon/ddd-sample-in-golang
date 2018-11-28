package service

import (
	"errors"
	"github.com/agiledragon/ddd-dci-sample-in-golang/transfer-money/domain/model"
	"sync"
)

type AccountService struct {
	repo model.AccountRepo
}

func (this *AccountService) Create(accountId string, amount uint, phoneNumber string) {
	account := model.AccountFactory{}.Create(accountId, amount, phoneNumber)
	this.repo.Add(account)
}

func (this *AccountService) TransferMoney(srcAccountId, dstAccountId string, amount uint) {
	from := this.repo.Get(srcAccountId)
	to := this.repo.Get(dstAccountId)
	from.MoneySource.TransferMoneyTo(to.MoneyDestination, amount)
}

func (this *AccountService) GetAmount(accountId string) (uint, error) {
	if account := this.repo.Get(accountId); account != nil {
		return account.MoneySource.GetAmount(), nil
	}
	return 0, errors.New("accountId is not existed")

}

func (this *AccountService) Destroy(accountId string) {
	this.repo.Remove(accountId)
}

var once sync.Once
var serviceInst *AccountService = &AccountService{}

func GetAccountService() *AccountService {
	once.Do(func() {
		serviceInst.repo = model.GetAccountRepo()
	})
	return serviceInst
}
