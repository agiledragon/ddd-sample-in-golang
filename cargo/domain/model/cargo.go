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

func (t *Cargo) Delay(days uint) {
	afterDays := t.GetAfterDays() + days
	t.Delivery = Delivery{AfterDays: afterDays}
}

func (t *Cargo) GetAfterDays() uint {
	return t.Delivery.AfterDays
}
