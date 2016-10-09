package volumeconv

import "fmt"

type Gallon float64
type Liter float64

func (gal Gallon) String() string { return fmt.Sprintf("%g gal", gal) }

func (l Liter) String() string { return fmt.Sprintf("%g L", l) }
