package lengthconv

import "fmt"

type Mile float64
type Kilometer float64

func (mi Mile) String() string { return fmt.Sprintf("%g mi", mi) }

func (km Kilometer) String() string { return fmt.Sprintf("%g km", km) }
