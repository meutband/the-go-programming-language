# Answer

When the help message gets generated, the custom `celciusFlag.String()` method is called, which contains `°C` in the string output. The default `celciusFlag` type is a custom `float64` value. 
