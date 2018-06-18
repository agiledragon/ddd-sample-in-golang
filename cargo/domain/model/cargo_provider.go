package model

type CargoProvider interface {
    Confirm(cargo *Cargo)
}

var p CargoProvider = nil

func SetCargoProvider(provider CargoProvider) {
    p = provider
}

func GetCargoProvider() CargoProvider {
    return p
}
