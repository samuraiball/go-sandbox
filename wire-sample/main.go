package main

import "github.com/samuraiball/go-sandbox/config"

func main() {
	shop := config.InitializeCakeShop()
	shop.Sell()
}
