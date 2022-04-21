package figuras

import (
	"math"
)

type Circulo struct {
	Radio float64
}

func (ci *Circulo) area() float64 {
	return math.Pi * (ci.Radio * ci.Radio)
}

func (ci *Circulo) perimetro() float64 {
	return 2 * math.Pi * ci.Radio
}
