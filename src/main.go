package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	mypkh "platzi/src/mypkg"
)

type car struct {
	brand string
	name  string
	year  int
	motor string
	color string
}

func isPalindromo(text string) bool {
	lenText := len(text)
	newText := strings.ToUpper(text)
	k := 0
	for j := lenText; j > 0; j-- {
		if newText[k] == newText[j-1] {
			k++
			continue
		} else {
			return false
		}
	}
	return true
}

func validateParImpar(number int) {
	if number%2 == 0 {
		fmt.Printf("El numero: %d es Par\n", number)
	} else {
		fmt.Printf("El numero: %d es Inpar\n", number)
	}
}

func login(username, password string) bool {
	return username == password && username == "debug"
}

func normalFuncion(message string) {
	fmt.Println(message)
}

func returnValue(a int) int {
	return a * 2
}

func tripleReturn(j int) (a, b, c int, d string) {
	return j * 1, j * 2, j * 3, "Hola"
}

func calcAreaRectangulo(baseRectangulo, alturaRectangulo float64) float64 {
	return baseRectangulo * alturaRectangulo
}

func calcAreaTrapecio(baseMenorTrapecio, baseMayorTrapecio, alturaTrapecio float64) float64 {
	return ((baseMenorTrapecio + baseMayorTrapecio) * alturaTrapecio) / 2
}

func calcAreaCiculo(radioCirculo float64) float64 {
	return math.Phi * math.Pow(float64(radioCirculo), 2)
}

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

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)

	normalFuncion("Hola")
	fmt.Println(returnValue(5))
	fmt.Println(tripleReturn(5))

	valueA, valueB, _, _ := tripleReturn(90)

	fmt.Println(valueA, valueB)

	resultA := calcAreaRectangulo(50, 100)
	resultB := calcAreaTrapecio(10, 20, 30)
	resultC := calcAreaCiculo(60)

	fmt.Printf("%f ---- %f --- %f", resultA, resultB, resultC)
	fmt.Println()

	for i := 0; i < 10; i++ {
		fmt.Println(" ->  ", i)
	}

	count := 0
	for count < 10 {
		fmt.Println(" ->  ", count)
		count++
	}

	// reto
	for z := 100; z > 0; z-- {
		fmt.Println(" ->  ", z)
	}

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es uno")
	} else {
		fmt.Println("No es uno")
	}

	// and
	if valor1 == 1 && valor2 == 1 {
		fmt.Println(" es verdad and")
	}

	if valor1 == 0 || valor2 != 0 {
		fmt.Print(" es verdad or ")
	}

	value, error := strconv.Atoi("100")

	if error != nil {
		fmt.Println()
		log.Fatal(error)
	} else {
		fmt.Println(" value - ", value)
	}

	validateParImpar(75)
	validateParImpar(76)

	fmt.Println(login("debug", "debug"))
	fmt.Println(login("admin", "admin"))

	modulo := 5 % 2

	switch {
	case modulo == 0:
		fmt.Println("Es par")
	case modulo != 0:
		fmt.Println("No es par")
	}

	//Defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// continue and break

	for w := 0; w < 10; w++ {
		fmt.Println(" -> ", w)

		if w == 2 {
			continue
		}
		fmt.Println(" -> ", w)
		if w == 7 {
			break
		}
	}

	var array_ [4]int
	fmt.Println(" -> ", array_)

	sliceOne := []int{1, 2, 3, 4}
	sliceOne[3] = 4
	fmt.Println("slice :  ", sliceOne, len(sliceOne), cap(sliceOne))

	fmt.Println(sliceOne[1])   // 2
	fmt.Println(sliceOne[1:3]) // 2,3
	fmt.Println(sliceOne[1:])  // 2,3,4
	fmt.Println(sliceOne[:2])  // 1,2

	sliceOne = append(sliceOne, 5)
	fmt.Println(" sliceOne ", sliceOne)

	newSlice := []int{100, 200, 300, 400}
	fmt.Println("newSlice ", newSlice)

	sliceOne = append(sliceOne, newSlice...)
	fmt.Println("sliceOne ", sliceOne)

	sliceString := []string{"hola", "que", "hace"}

	for i, value := range sliceString {
		fmt.Println(i, value)
	}
	fmt.Println("IsPali: ", isPalindromo("Anita Lava La Tina"))

	mapa := make(map[string]int)
	mapa["Richi"] = 33
	mapa["Ivon"] = 31
	mapa["Marco"] = 29

	fmt.Println(mapa)

	for i, v := range mapa {
		fmt.Println(i, v)
	}

	valorE, exist := mapa["Ivonn"]

	fmt.Println(valorE, exist)

	myCar := car{name: "bocho", year: 1991}
	fmt.Println(myCar)

	var myCar2 car
	myCar2.name = "mustang"

	fmt.Println(myCar2)

	var myCar3 mypkh.CarPublic
	myCar3.Name = "Bocho"
	fmt.Println(myCar3)

	mypkh.PrintMessage()

	a1 := 10
	b1 := &a1
	*b1 = 500

	fmt.Println(a1, b1, *b1)

	myPc := mypkh.Pc{Ram: 64, Disk: 1, Cpu: "intel"}
	fmt.Println(myPc)

	myPc.Ping()

	fmt.Println(myPc)

	myPc.DuplicateRam()
	myPc.DuplicateRam()

	fmt.Println(myPc)

	myPc1 := mypkh.Pc{Ram: 64, Disk: 1, Cpu: "amd"}
	fmt.Println(myPc1)
	fmt.Println(myPc1.String())

}
