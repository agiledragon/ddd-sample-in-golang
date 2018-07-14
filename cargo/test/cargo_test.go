package test

import (
	"github.com/agiledragon/ddd-sample-in-golang/cargo/app/service"
	"github.com/agiledragon/ddd-sample-in-golang/cargo/domain/model"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type SpyCargoProvider struct {
	cargoId   string
	afterDays uint
}

func (this *SpyCargoProvider) Confirm(cargo *model.Cargo) {
	this.cargoId = cargo.Id()
	this.afterDays = cargo.GetAfterDays()
}

type FakeCargoRepo struct {
	cargoes map[string]*model.Cargo
}

func (this *FakeCargoRepo) Add(cargo *model.Cargo) {
	this.cargoes[cargo.Id()] = cargo
}

func (this *FakeCargoRepo) Get(cargoId string) *model.Cargo {
	return this.cargoes[cargoId]
}

func (this *FakeCargoRepo) Update(cargo *model.Cargo) {
	this.cargoes[cargo.Id()] = cargo
}

func (this *FakeCargoRepo) Remove(cargoId string) {
	delete(this.cargoes, cargoId)
}

func TestCargo(t *testing.T) {
	provider := &SpyCargoProvider{}
	model.SetCargoProvider(provider)
	repo := &FakeCargoRepo{make(map[string]*model.Cargo)}
	model.SetCargoRepo(repo)
	const cargoId = "1"

	Convey("TestCargo", t, func() {

		Convey("create cargo", func() {
			const afterDays = 10
			service.CreateCargo(cargoId, afterDays)
			So(provider.cargoId, ShouldEqual, cargoId)
			So(provider.afterDays, ShouldEqual, afterDays)
			So(service.GetCargoAfterDays(cargoId), ShouldEqual, afterDays)
			service.DestroyCargo(cargoId)
		})

		Convey("delay cargo", func() {
			const afterDays = 20
			const days = 5
			service.CreateCargo(cargoId, afterDays)
			service.DelayCargo(cargoId, days)
			So(provider.cargoId, ShouldEqual, cargoId)
			So(provider.afterDays, ShouldEqual, afterDays+days)
			So(service.GetCargoAfterDays(cargoId), ShouldEqual, afterDays+days)
			service.DestroyCargo(cargoId)
		})
	})
}
