package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	//with or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad, OK")
	}

	//Convertir texto a numero
	value, err := strconv.Atoi("53")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Value:", value)

}
