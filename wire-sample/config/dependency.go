// +build wireinject

package config

import (
	"github.com/google/wire"
	"github.com/samuraiball/go-sandbox/message"
	"github.com/samuraiball/go-sandbox/sweets"
)

func InitializeEvent() message.Event {

	wire.Build(
		message.NewEvent,
		message.NewGreeter,
		message.NewMassage)

	return message.Event{}
}

func InitializeCakeShop() sweets.CakeShop {
	wire.Build(
		sweets.CakeShopWithMontBlancProvider,
		sweets.MontBlancProvider,
		sweets.ChestnutProvider,
	)
	return sweets.CakeShop{}
}
