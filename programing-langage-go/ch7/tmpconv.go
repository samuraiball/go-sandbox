package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "℃":
		f.Celsius = Celsius(value)
		return nil
	case "F", "℉":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperture %q", s)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g ℃", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g °F", f)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func CToF(celsius Celsius) Fahrenheit {
	return Fahrenheit(celsius*9/5 + 32)
}

func FToC(fahrenheit Fahrenheit) Celsius {
	return Celsius((fahrenheit - 32) * 5 / 9)
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
