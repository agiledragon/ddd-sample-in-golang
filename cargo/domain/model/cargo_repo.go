package model

type CargoRepo interface {
	Add(cargo *Cargo)
	Get(cargoId string) *Cargo
	Update(cargo *Cargo)
	Remove(cargoId string)
}

var r CargoRepo = nil

func SetCargoRepo(repo CargoRepo) {
	r = repo
}

func GetCargoRepo() CargoRepo {
	return r
}
