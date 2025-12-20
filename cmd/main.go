package main

import (
	"fmt"

	"github.com/viktorHadz/webserver-go/interfaces"
)

func main() {
	fmt.Printf("--- Main Start ---\n\n")

	shapes := []interfaces.Shape{
		interfaces.NewCircle(6),
		interfaces.NewCircle(8),
		interfaces.NewCircle(10),
		interfaces.NewSquare(10),
		interfaces.NewSquare(20),
		interfaces.NewSquare(30),
	}
	for _, s := range shapes {
		interfaces.PrintShape(s)
	}

	dog := interfaces.Dog{Name: "Ugly"}
	person := interfaces.Person{Name: "Orpheus"}
	
	dog.Speak("WOOF WOOF MOTHERFUCKER")
	dog.Speak("WOOF WOOF 2")
	person.Speak("Yo! Wassup")
}
