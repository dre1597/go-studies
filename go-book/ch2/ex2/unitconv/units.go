package unitconv

import "fmt"

type Pascal float64
type Atmospheric float64

func (pa Pascal) String() string       { return fmt.Sprintf("%g Pa", pa) }
func (atm Atmospheric) String() string { return fmt.Sprintf("%g atm", atm) }
