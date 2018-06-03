package ft

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
    "github.com/agiledragon/ddd-sample-in-golang/domain/model"
    "github.com/agiledragon/ddd-sample-in-golang/app/service"
)

type SpyCargoProvider struct {
    cargoId string
    afterDays uint
}

func (t *SpyCargoProvider) Confirm(cargo *model.Cargo) {
    t.cargoId = cargo.GetId()
    t.afterDays = cargo.GetAfterDays()
}

type FakeCargoRepo struct {
    cargoes map[string]*model.Cargo
}

func (t *FakeCargoRepo) Add(cargo *model.Cargo) {
    t.cargoes[cargo.GetId()] = cargo
}

func (t *FakeCargoRepo) Get(cargoId string) *model.Cargo {
    return t.cargoes[cargoId]
}

func (t *FakeCargoRepo) Update(cargo *model.Cargo) {
    t.cargoes[cargo.GetId()] = cargo
}

func (t *FakeCargoRepo) Remove(cargoId string) {
    delete(t.cargoes, cargoId)
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
        })

        Convey("delay cargo", func() {
            const afterDays = 20
            const days = 5
            service.CreateCargo(cargoId, afterDays)
            service.DelayCargo(cargoId, days)
            So(provider.cargoId, ShouldEqual, cargoId)
            So(provider.afterDays, ShouldEqual, afterDays + days)
            So(service.GetCargoAfterDays(cargoId), ShouldEqual, afterDays + days)

        })
    })
}

