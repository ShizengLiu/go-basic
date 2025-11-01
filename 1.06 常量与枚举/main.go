package main

import (
	"fmt"
)

const mainName string = "main"

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

func method(Gender Gender) Gender {
	return Gender
}

func (g *Gender) GetGenderString() string {
	switch *g {
	case Male:
		return string(Male)
	case Female:
		return string(Female)
	default:
		return "Unknown"
	}
}

func (g *Gender) IsMale() bool {
	return *g == Male

}

func (g *Gender) changeGender() {
	switch *g {
	case Male:
		*g = Female
	case Female:
		*g = Male
	default:
		*g = "Unknown"
	}
}

type Person byte

const (
	MaleP Person = iota
	FemaleP
)

func main() {
	fmt.Println("main method invoked!")
	fmt.Println(mainName)
	const mainFuncName = "mainFunc"
	fmt.Println(mainFuncName)

	fmt.Println(method(Male))

	var person Gender = Female
	fmt.Println(person.GetGenderString())
	fmt.Println(person.IsMale())
	person.changeGender()
	fmt.Println("------------")
	fmt.Println(person.GetGenderString())
	fmt.Println(person.IsMale())

	var p Person = MaleP
	fmt.Println(p)
	var p1 Person = FemaleP
	fmt.Println(p1)

}
