package figuras

import "fmt"

type figurasG interface {
	area() float64
	perimetro() float64
}

func CalcularFiguras(fig figurasG) {
	fmt.Println("Medidas", fig)
	fmt.Println("Area: ", fig.area())
	fmt.Println("Perimetro: ", fig.perimetro())
}
