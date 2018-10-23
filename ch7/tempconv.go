package tempconv

import ("fmt"
        "flag"
      )

type Celsius float64
type Farenheit float64

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c)}
func (f Farenheit) String() string { return fmt.Sprintf("%g°F", f)}

// CToF converts Celsius temp to Farenheit
func CToF(c Celsius) Farenheit { return Farenheit(c*9/5 + 32)}
// FToC converts Farenheit temp to Celsius
func FToC(f Farenheit) Celsius { return Celsius((f-32) * 5/9)}
type celsiusFlag struct { Celsius}

func (f *celsiusFlag) Set(s string) error {
  var unit string
  var value float64
  fmt.SScanf(s, "%f%s", &value, &unit) // no error check needed
  switch unit {
  case "C", "°C" :
    f.Celsius = Celsius(value)
    return nil
  case "F", "°F" :
    f.Celsius = FToC(Farenheit(value))
    return nil
    }
  return fmt.Errorf("invalid temperature %q", s)
}
// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag
// variable
// The flag argument must have a quantity and a unit, e.g. "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
  f := celsiusFlag(value)
  flag.CommandLine.Var(&f, name, usage)
  return &f.Celsius
}
