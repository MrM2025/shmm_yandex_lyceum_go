package main

import "fmt"

type Animal interface {
	MakeSound() string
}

type Dog struct {
	sound string
}

type Cat struct {
	sound string
}

func (d Dog) MakeSound() {
	fmt.Println("Гав!")
}

func (c Cat) MakeSound() {
	fmt.Println("Мяу!")
}

func main() {
	cat1 := Cat{}
	cat1.MakeSound()
	dog1 := Dog{}
	dog1.MakeSound()
}
