package main

import "fmt"

func main() {

	//For condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("================\n")

	//For while
	count := 0
	for count < 10 {
		fmt.Println(count)
		count++
	}

	fmt.Printf("================\n")

	//For forever
	counfForever := 0
	for {
		fmt.Println(counfForever)
		counfForever++
	}

}
