package main

import "fmt"

func main() {
	//Defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//continue and break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//Continue
		if i == 2 {
			fmt.Println("es 2")
			continue
		}

		//break
		if i == 8 {
			fmt.Println("Break")
			break
		}

	}
}
