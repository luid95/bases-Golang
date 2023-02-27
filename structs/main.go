package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {

	//Instanciar mi struct
	myCar := car{brand: "Ford", year: 2020}

	fmt.Println(myCar)

	//Otra manera
	var otherCar car

	otherCar.brand = "Ferrari"

	fmt.Println(otherCar)

}
