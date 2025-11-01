package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name  string `json:"name"`
	Age   int
	Man   bool
	Call  func() byte
	Map   map[string]string
	Ch    chan string
	Arr   [32]uint8
	Slice []interface{}
	Ptr   *int
	once  sync.Once
}

type Custom struct {
	int
	string
	Other string
}

type Animal struct {
	typeName string
}

func (a Animal) string() string {
	return a.typeName
}

func (a Animal) stringA() string {
	return a.typeName
}

func (a Animal) setTypeName(v string) {
	a.typeName = v
}

func (a *Animal) stringPTypeName() string {
	return a.typeName
}

func (a *Animal) setPTypeName(v string) {
	a.typeName = v
}

type Human struct {
	Animal
	name string
}

func (b Human) string() string {
	return b.name
}

func (b Human) stringB() string {
	return b.name
}

type Woman struct {
	Human
	typeName string
	name     string
	fat      bool
	weight   []byte
}

func (c Woman) bool() bool {
	return c.fat
}

func (c Woman) modityWeight() {
	c.weight[2] = 100
}

func callStructMethod() {
	var a Animal
	a = Animal{
		typeName: "person",
	}
	a.string()
}

func NewWoman() Woman {
	return Woman{
		Human: Human{
			Animal: Animal{
				typeName: "person1",
			},
			name: "good woman",
		},
		typeName: "person2",
		name:     "good woman2",
		fat:      false,
		weight:   []byte{100, 50, 100},
	}
}

func main() {
	fmt.Println("main method invoked!")
	c := NewWoman()
	cp := &c
	fmt.Println(c.string())
	fmt.Println(c.stringA())
	fmt.Println(c.stringB())

	fmt.Println(cp.string())
	fmt.Println(cp.stringA())
	fmt.Println(cp.stringB())

	c.setTypeName("1woman")
	fmt.Println("--------c.setType")
	fmt.Println(c.Animal.typeName)
	fmt.Println(cp.Animal.typeName)

	c.setTypeName("1woman")
	fmt.Println("--------c.setType")
	fmt.Println(c.Animal.typeName)
	fmt.Println(cp.Animal.typeName)

	cp.setTypeName("2woman")
	fmt.Println("--------cp.setType")
	fmt.Println(c.Animal.typeName)
	fmt.Println(cp.Animal.typeName)

	c.setPTypeName("3woman")
	fmt.Println("--------c.setPType")
	fmt.Println(c.Animal.typeName)
	fmt.Println(cp.Animal.typeName)

	cp.setPTypeName("4woman")
	fmt.Println("--------cp.setPType")
	fmt.Println(c.Animal.typeName)
	fmt.Println(cp.Animal.typeName)

	cp.modityWeight()
	fmt.Println("--------cp.modityWeight")
	fmt.Println(c.weight)
	fmt.Println(cp.weight)
}
