package sweets

import "fmt"

type Chestnut struct {
}

type Cake interface {
	Eaten()
}

type MontBlanc struct {
	Chestnut Chestnut
}

func (m MontBlanc) Eaten() {
	fmt.Printf("I'm MontBlanc, How is it?")
}

type CakeShop struct {
	Cake Cake
}

func (c CakeShop) Sell() {
	c.Cake.Eaten()
}

func MontBlancProvider(c Chestnut) MontBlanc {
	return MontBlanc{Chestnut: c}
}

func ChestnutProvider() Chestnut {
	return Chestnut{}
}

func CakeShopWithMontBlancProvider(c MontBlanc) CakeShop {
	return CakeShop{Cake: c}
}
