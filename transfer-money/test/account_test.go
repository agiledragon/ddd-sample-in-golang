package test

import (
	"fmt"
	"github.com/agiledragon/ddd-dci-sample-in-golang/transfer-money/app/service"
	"github.com/agiledragon/ddd-dci-sample-in-golang/transfer-money/domain/model"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type SpyPhoneProvider struct {
	phoneNumbers []string
	msgs         []string
}

func (this *SpyPhoneProvider) Send(phoneNumber, msg string) {
	this.phoneNumbers = append(this.phoneNumbers, phoneNumber)
	this.msgs = append(this.msgs, msg)
}

type FakeAccountRepo struct {
	accounts map[string]*model.Account
}

func (a *FakeAccountRepo) Add(account *model.Account) {
	a.accounts[account.Id()] = account
}

func (a *FakeAccountRepo) Get(accountId string) *model.Account {
	return a.accounts[accountId]
}

func (a *FakeAccountRepo) Remove(accountId string) {
	delete(a.accounts, accountId)
}

func TestAccount(t *testing.T) {
	provider := &SpyPhoneProvider{
		phoneNumbers: make([]string, 0),
		msgs:         make([]string, 0),
	}
	model.SetPhoneProvider(provider)
	repo := &FakeAccountRepo{make(map[string]*model.Account)}
	model.SetAccountRepo(repo)

	Convey("TestAccount", t, func() {

		Convey("create account", func() {
			accountId := "zhangsan123"
			initialAmount := uint(1000)
			phoneNumber := "13310998098"
			service.CreateAccount(accountId, initialAmount, phoneNumber)
			actualAmount, err := service.GetAmount(accountId)
			So(err, ShouldBeNil)
			So(actualAmount, ShouldEqual, initialAmount)
			service.DestroyAccount(accountId)
		})

		Convey("transfer money", func() {
			srcAccountId := "zhangsan123"
			srcInitialAmount := uint(1000)
			srcPhoneNumber := "13310998098"
			service.CreateAccount(srcAccountId, srcInitialAmount, srcPhoneNumber)

			dstAccountId := "lisi456"
			dstInitialAmount := uint(200)
			dstPhoneNumber := "19916588070"
			service.CreateAccount(dstAccountId, dstInitialAmount, dstPhoneNumber)

			amount := uint(300)
			service.TransferMoney(srcAccountId, dstAccountId, amount)
			srcActualAmount, err := service.GetAmount(srcAccountId)
			So(err, ShouldBeNil)
			So(srcActualAmount, ShouldEqual, srcInitialAmount-amount)
			dstActualAmount, err := service.GetAmount(dstAccountId)
			So(err, ShouldBeNil)
			So(dstActualAmount, ShouldEqual, dstInitialAmount+amount)

			srcMsg := fmt.Sprintf("%s send %d money to %s", srcAccountId, amount, dstAccountId)
			dstMsg := fmt.Sprintf("%s receive %d money from %s", dstAccountId, amount, srcAccountId)
			So(provider.phoneNumbers[0], ShouldEqual, dstPhoneNumber)
			So(provider.msgs[0], ShouldEqual, dstMsg)
			So(provider.phoneNumbers[1], ShouldEqual, srcPhoneNumber)
			So(provider.msgs[1], ShouldEqual, srcMsg)

			service.DestroyAccount(srcAccountId)
			service.DestroyAccount(dstAccountId)
		})
	})
}
