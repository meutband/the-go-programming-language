// This file performs Pounds and Kilograms conversions.
// This file was added for Exercise 2.2.
package tempconv

import "fmt"

type Pound float64
type Kilogram float64

func (p Pound) String() string     { return fmt.Sprintf("%g lbs", p) }
func (kg Kilogram) String() string { return fmt.Sprintf("%g kg", kg) }
