package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*(9/5) + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(KToC(k))) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return Kelvin(CToK(FToC(f))) }

// The following were added for Exercise 2.2

// FToM converts a Feet length to Meters
func FToM(f Feet) Meter { return Meter(f * 0.3048) }

// MToF converts a Meters length to Feet
func MToF(m Meter) Feet { return Feet(m * 3.28084) }

// PToKg converts a Pound length to Kilogram
func PToKg(p Pound) Kilogram { return Kilogram(p * 0.453592) }

// KgToP converts a Kilogram length to Pound
func KgToP(kg Kilogram) Pound { return Pound(kg * 2.20462) }
