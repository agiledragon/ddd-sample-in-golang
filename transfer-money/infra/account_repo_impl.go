package infra

import "github.com/agiledragon/ddd-dci-sample-in-golang/transfer-money/domain/model"

type AccountRepoImpl struct {
}

func (this *AccountRepoImpl) Add(account *model.Account) {

}

func (this *AccountRepoImpl) Get(accountId string) *model.Account {
	return nil
}

func (this *AccountRepoImpl) Remove(accountId string) {

}
