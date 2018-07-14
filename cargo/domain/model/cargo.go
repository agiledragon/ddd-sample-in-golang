package model

import (
	"github.com/agiledragon/ddd-sample-in-golang/cargo/domain/model/base"
)

type Cargo struct {
	base.AggregateRoot
	Delivery Delivery
}

func newCargo(cargoId string, delivery Delivery) *Cargo {
	return &Cargo{AggregateRoot: base.NewAggregateRoot(cargoId), Delivery: delivery}
}

func (this *Cargo) Delay(days uint) {
	afterDays := this.GetAfterDays() + days
	this.Delivery = Delivery{AfterDays: afterDays}
}

func (this *Cargo) GetAfterDays() uint {
	return this.Delivery.AfterDays
}
