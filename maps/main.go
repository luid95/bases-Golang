package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepe"] = 20

	fmt.Println(m)

	//Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar valor
	value, ok := m["Josep"]
	fmt.Println(value, ok)

}
