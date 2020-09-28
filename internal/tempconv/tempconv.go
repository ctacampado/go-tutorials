package tempconv

import "fmt"

// Celsius type
type Celsius float64

// Fahrenheit type
type Fahrenheit float64

const (
	// AbsoluteZeroC temp in Celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingC temp in Celsius
	FreezingC Celsius = 0
	// BoilingC temp in Celsius
	BoilingC Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%gC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gF", f) }
