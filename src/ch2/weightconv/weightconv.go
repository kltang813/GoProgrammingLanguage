package weightconv

import "fmt"

type Pound float64
type Kilogram float64

func (p Pound) String() string {
	return fmt.Sprintf("%g lbs", p)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%g kg", k)
}
