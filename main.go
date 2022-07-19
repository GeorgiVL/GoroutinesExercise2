package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	First  string
	Second string
	Age    int
}

func (p *person) speak() {
	fmt.Printf("%s\t%s is speaking right now", p.First, p.Second)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		First:  "Georgi",
		Second: "Valkanov",
		Age:    24,
	}

	saySomething(&p1)
	fmt.Println("This code was added to github")
}
