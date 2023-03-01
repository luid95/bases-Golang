package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2) //numero de chanels

	c <- "mensaje1"
	c <- "mensaje2"

	fmt.Println(len(c), cap(c))

	//Range y close
	//close -> dicir que va a cerrar el canal
	close(c)

	// c <- "mensaje3"
	//Rage -> recorrido de una lista

	for message := range c {
		fmt.Println(message)
	}

	//Select
	email1 := make(chan string) //declarado sin limita de capacidad
	email2 := make(chan string)

	go message("mensaje 1", email1)
	go message("mensaje 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido", m2)
		}
	}

}
