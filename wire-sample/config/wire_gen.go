// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package config

import (
	"github.com/samuraiball/go-sandbox/message"
	"github.com/samuraiball/go-sandbox/sweets"
)

// Injectors from dependency.go:

func InitializeEvent() message.Event {
	messageMessage := message.NewMassage()
	greeter := message.NewGreeter(messageMessage)
	event := message.NewEvent(greeter)
	return event
}

func InitializeCakeShop() sweets.CakeShop {
	chestnut := sweets.ChestnutProvider()
	montBlanc := sweets.MontBlancProvider(chestnut)
	cakeShop := sweets.CakeShopWithMontBlancProvider(montBlanc)
	return cakeShop
}
