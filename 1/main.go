package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Greeting() {
	fmt.Printf("Привет, меня зовут %v, мне %v\n", h.Name, h.Age)
}

type Action struct {
	Human
	Job string
}

func (a *Action) DoWork() {
	fmt.Printf("%v %v начал свою работу...\n", a.Job, a.Name)
}

func main() {
	human := Action{
		Human: Human{Name: "Aleksej",
			Age: 35,
		},
		Job: "Go developer",
	}
	human.Greeting()
	human.DoWork()
}
