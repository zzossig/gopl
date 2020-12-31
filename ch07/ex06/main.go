// Add support for Kelvin temperatures to `tempflag`.

package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }
func FToK(f Fahrenheit) Kelvin { return Kelvin((f - 32) * 5/9 + 273.15) }
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k - 273.15) * 9/5 + 32) }

type celsiusFlag struct{ Celsius }

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(temp)
}

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (k Kelvin) String() string { return fmt.Sprintf("%gK", k) }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}