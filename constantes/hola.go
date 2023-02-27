package main

import "fmt"

func normalFunction(message string) {

	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {

	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println(pi)
	fmt.Println(pi2)

	//Declaracion de variable enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area del cuadrado

	const baseCuadrado = 10

	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area del cuadrado:", areaCuadrado)

	x := 10
	y := 50

	//Suma
	result := x + y

	fmt.Println("Suma:", result)

	//Resta
	result2 := y - x

	fmt.Println("Resta:", result2)

	//Multiplicacion
	result3 := y * x

	fmt.Println("Resta:", result3)

	//Division
	result4 := y / x

	fmt.Println("Resta:", result4)

	//modulo
	result5 := y % x

	fmt.Println("Residuo:", result5)

	//incremental
	x++

	fmt.Println("Incremental:", x)

	//decremental
	x--

	fmt.Println("Decremental:", x)

	//Declaracion de variables
	helloMessage := "Hello"
	worlMessage := "world"

	//Println
	fmt.Println(helloMessage, worlMessage)

	//Printf
	nombre := "Platzi"
	cursos := 500

	fmt.Printf("%s tiene mas de %d cursos \n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos \n", nombre, cursos)

	//Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos \n", nombre, cursos)

	fmt.Println(message)

	//Tipo de dato
	fmt.Printf("Tipo de dato de helloMessage: %T\n", helloMessage)
	fmt.Printf("Tipo de dato de cursos: %T\n", cursos)

	normalFunction("Hola mundo")
	tripleArgument(1, 2, "Hola")

	value := returnValue(2)
	fmt.Println("Value: ", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("Value: ", value1, value2)
}
