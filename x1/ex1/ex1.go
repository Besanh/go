package main

import (
	"code/ex2"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("Prefix loi: ")
	log.SetFlags(0)

	names := []string{"An", "Thu", "Bi"}
	message, error := ex2.Ex3(names)
	if error != nil {
		log.Fatalln(error)
	}
	fmt.Println(message)
}
