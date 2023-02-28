package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
	fmt.Println(myPC.brand, "duplicateRAM")
}

func main() {

	a := 50
	b := &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPc := pc{ram: 16, disk: 20, brand: "msi"}
	fmt.Println(myPc)

	myPc.ping()

	fmt.Println(myPc)
	myPc.duplicateRAM()

	fmt.Println(myPc)
	myPc.duplicateRAM()

	fmt.Println(myPc)

}
