package service

import (
    "github.com/agiledragon/ddd-sample-in-golang/cargo/domain/service"
)

func CreateCargo(cargoId string, afterDays uint) {
    service.GetCargoServiceInstance().Create(cargoId, afterDays)
}

func DelayCargo(cargoId string, days uint) {
    service.GetCargoServiceInstance().Delay(cargoId, days)
}

func GetCargoAfterDays(cargoId string) uint {
    return service.GetCargoServiceInstance().GetAfterDays(cargoId)
}