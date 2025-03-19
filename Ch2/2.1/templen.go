// This file performs Feet and Length conversions.
// This file was added for Exercise 2.2.
package tempconv

import "fmt"

type Feet float64
type Meter float64

func (f Feet) String() string  { return fmt.Sprintf("%g ft", f) }
func (m Meter) String() string { return fmt.Sprintf("%g m", m) }
