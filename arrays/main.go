package main

import (
	"fmt"
	"strings"
)

func isPalindormo(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if strings.ToLower(text) == strings.ToLower(textReverse) {
		fmt.Println("Es palindrmo")
	} else {
		fmt.Println("No es palindrmo")

	}
}

func main() {
	//Array
	var array [4]int
	array[0] = 1
	array[1] = 2

	fmt.Println(array, len(array), cap(array))

	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println(slice, len(slice), cap(slice))

	//Metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//Append
	slice = append(slice, 7)

	fmt.Println(slice)

	//Append nueva list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)

	fmt.Println(slice)

	fmt.Println("Recorrido de Slices")
	//Recorrido de slice
	mySlice := []string{"Hola", "que", "hace"}

	fmt.Println("Obtener solo el valor")
	for _, valor := range mySlice {
		fmt.Println(valor)

	}

	fmt.Println("Obtener solo el indice")
	for i := range mySlice {
		fmt.Println(i)

	}

	fmt.Println("Obtener solo el indice y valor")
	for i, valor := range mySlice {
		fmt.Println(i, valor)

	}

	//Palindromo
	isPalindormo("Ama")
}
