package main

import (
	"fmt"

	pk "./mypackage"
)

func main() {
	var myCar = pk.CarPublic
	myCar.Brand = "ferrari"

	fmt.Println(myCar)
}
