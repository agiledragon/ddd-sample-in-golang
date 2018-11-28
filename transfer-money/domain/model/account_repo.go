package model

var repoInst AccountRepo

type AccountRepo interface {
	Add(account *Account)
	Get(accountId string) *Account
	Remove(accountId string)
}

func SetAccountRepo(repo AccountRepo) {
	repoInst = repo
}

func GetAccountRepo() AccountRepo {
	return repoInst
}
