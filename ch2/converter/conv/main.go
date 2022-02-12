package conv

import "fmt"

type Foot float64
type Meter float64

type Kilogram float64
type Pound float64

func (f Foot) String() string     { return fmt.Sprintf("%g ft", f) }
func (m Meter) String() string    { return fmt.Sprintf("%g m", m) }
func (k Kilogram) String() string { return fmt.Sprintf("%g kg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%g lb", p) }

func PToKg(p Pound) Kilogram {
	return Kilogram(p) * 0.453592
}

func KgToP(k Kilogram) Pound {
	return Pound(k) * 2.20462
}

func FtToM(f Foot) Meter {
	return Meter(f) * 0.3048
}

func MtoFt(m Meter) Foot {
	return Foot(m) * 3.281
}
