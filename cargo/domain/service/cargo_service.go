package service

import (
	"github.com/agiledragon/ddd-sample-in-golang/cargo/domain/model"
	"sync"
)

type CargoService struct {
	repo     model.CargoRepo
	provider model.CargoProvider
}

var cs = &CargoService{}
var once sync.Once

func GetCargoService() *CargoService {
	once.Do(func() {
		cs.repo = model.GetCargoRepo()
		cs.provider = model.GetCargoProvider()
	})
	return cs
}

func (this *CargoService) Create(cargoId string, afterDays uint) {
	cargo := model.CargoFactory{}.Create(cargoId, afterDays)
	this.repo.Add(cargo)
	this.provider.Confirm(cargo)
}

func (this *CargoService) Delay(cargoId string, days uint) {
	cargo := this.repo.Get(cargoId)
	if cargo == nil {
		panic("not found cargo by cargoId")
	}
	cargo.Delay(days)
	this.repo.Update(cargo)
	this.provider.Confirm(cargo)
}

func (this *CargoService) GetAfterDays(cargoId string) uint {
	cargo := this.repo.Get(cargoId)
	if cargo == nil {
		panic("not found cargo by cargoId")
	}
	return cargo.GetAfterDays()
}

func (this *CargoService) DestroyCargo(cargoId string) {
	this.repo.Remove(cargoId)
}
