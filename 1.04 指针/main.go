package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("main method invoked!")
	var p1 *int
	var p2 *string

	i := 1
	s := "string"

	p1 = &i
	p2 = &s

	p3 := &p2
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	fmt.Println(*p1 == i)
	*p1 = 2
	fmt.Println(i)
	fmt.Println("----------------------")
	a := 2
	var p *int
	fmt.Println(&a)
	p = &a
	fmt.Println(p)
	fmt.Println(&a)

	var pp **int
	pp = &p
	fmt.Println(pp, p)
	**pp = 3
	fmt.Println(pp, *pp, p)
	fmt.Println(**pp, *p)
	fmt.Println(a, &a)

	var ppp ***int
	ppp = &pp
	fmt.Println(ppp, pp, p)
	***ppp = 4
	fmt.Println(***ppp, **pp, *p)
	fmt.Println(a, &a)

	uua := "hello"

	upA := uintptr(unsafe.Pointer(&uua))
	fmt.Println(upA)
	upA += 1
	fmt.Println(upA)
	upAPtr := (*uint8)(unsafe.Pointer(upA))
	fmt.Println(*upAPtr)

}
