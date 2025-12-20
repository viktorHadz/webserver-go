package interfaces

import "fmt"

type Speaker interface {
	Speak()
}

type Person struct {
	Name string
}
type Dog struct {
	Name string
}

func (p Person) Speak(what string) {
	fmt.Println(what)
}
func (d Dog) Speak(what string) {
	fmt.Println(what)
}

func SaySomething(s Speaker, what string) {
	s.Speak()
}
