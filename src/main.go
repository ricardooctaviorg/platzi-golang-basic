package main

import (
	"fmt"
	"math"
)

func main() {
	const pi float64 = 3.1416
	const pi2 = 3.1416
	fmt.Println("Hola Mundo", pi)
	fmt.Println("Hola Mundo", pi2)

	var base int = 200
	var altura int = 300

	var area = (base * altura) / 2
	area2 := 500

	fmt.Println("Base: ", base)
	fmt.Println("Altura: ", altura)
	fmt.Println("Area: ", area)
	fmt.Println("Area2: ", area2)

	var a int
	var b string
	var c float64
	var d bool

	fmt.Println(a, b, c, d)

	// Area: Recatangulo, Trapecio y  Circulo

	baseRectangulo := 10
	alturaRectangulo := 20

	fmt.Println("Area rectangulo: ", baseRectangulo*alturaRectangulo)

	baseMenorTrapecio := 10
	baseMayorTrapecio := 30
	alturaTrapecio := 50

	areaTrapecio := ((baseMenorTrapecio + baseMayorTrapecio) * alturaTrapecio) / 2

	fmt.Println("Area trapecio: ", areaTrapecio)

	radioCirculo := 10

	areaCirculo := math.Phi * math.Pow(float64(radioCirculo), 2)

	fmt.Println("Area circulo : ", areaCirculo)

}
