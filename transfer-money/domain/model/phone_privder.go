package model

type PhoneProvider interface {
	Send(phoneNumber, msg string)
}

var providerInst PhoneProvider

func SetPhoneProvider(provider PhoneProvider) {
	providerInst = provider
}

func GetPhoneProvider() PhoneProvider {
	return providerInst
}
