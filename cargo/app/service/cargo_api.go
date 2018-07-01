package service

import (
    "github.com/agiledragon/ddd-sample-in-golang/cargo/domain/service"
)

func CreateCargo(cargoId string, afterDays uint) {
    service.GetCargoService().Create(cargoId, afterDays)
}

func DelayCargo(cargoId string, days uint) {
    service.GetCargoService().Delay(cargoId, days)
}

func GetCargoAfterDays(cargoId string) uint {
    return service.GetCargoService().GetAfterDays(cargoId)
}

func DestroyCargo(cargoId string) {
    service.GetCargoService().DestroyCargo(cargoId)
}
