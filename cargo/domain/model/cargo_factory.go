package model

type CargoFactory struct {
}

func (t CargoFactory) Create(cargoId string, afterDays uint) *Cargo {
	delivery := Delivery{AfterDays: afterDays}
	return newCargo(cargoId, delivery)
}
