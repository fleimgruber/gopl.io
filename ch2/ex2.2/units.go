package main

import (
	"fmt"
	"os"
	"strconv"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Meter float64
type Feet float64
type Pound float64
type Kilogram float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = 0
	FreezingK     Kelvin  = 273.15
	BoilingK      Kelvin  = 373.15
	FtToMRatio            = 0.3048
	PToKgRatio            = 0.45359237
)

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// KToC converts a Celsius temperature to Kelvin.
func KToC(k Kelvin) Celsius { return Celsius(k - FreezingK) }

// MToFt converts a metric length to imperial feet.
func MToFt(m Meter) Feet { return Feet(m / FtToMRatio ) }

// FtToM converts imperial feet to metric length.
func FtToM(ft Feet) Meter { return Meter(FtToMRatio * ft) }

// KgToP converts a Kilogram weight to imperial pounds.
func KgToP(kg Kilogram) Pound { return Pound(kg / PToKgRatio ) }

// PToKg converts imperial pounds to kilogram weight.
func PToKg(p Pound) Kilogram { return Kilogram(p * PToKgRatio ) }

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g K", k) }
func (m Meter) String() string      { return fmt.Sprintf("%g m", m) }
func (ft Feet) String() string      { return fmt.Sprintf("%g ft", ft) }
func (p Pound) String() string      { return fmt.Sprintf("%g lb", p) }
func (kg Kilogram) String() string      { return fmt.Sprintf("%g kg", kg) }

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "units: %v\n", err)
			os.Exit(1)
		}
		f := Fahrenheit(t)
		c := Celsius(t)
		k := Kelvin(t)
		m := Meter(t)
		ft := Feet(t)
		kg := Kilogram(t)
		p := Pound(t)
		fmt.Printf("%s = %s\n", f, FToC(f))
		fmt.Printf("%s = %s\n", c, CToF(c))
		fmt.Printf("%s = %s\n", k, KToC(k))
		fmt.Printf("%s = %s\n", m, MToFt(m))
		fmt.Printf("%s = %s\n", ft, FtToM(ft))
		fmt.Printf("%s = %s\n", kg, KgToP(kg))
		fmt.Printf("%s = %s\n", p, PToKg(p))
	}
}
