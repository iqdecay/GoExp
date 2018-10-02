package tempconv

// CToF converts Celsius temp to Farenheit
func CToF(c Celsius) Farenheit { return Farenheit(c*9/5 + 32)}
// FToC converts Farenheit temp to Celsius
func FToC(f Farenheit) Celsius { return Celsius((f-32) * 5/9)}

