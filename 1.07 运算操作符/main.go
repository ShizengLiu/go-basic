package main

import (
	"fmt"
)

func calOperation() {
	a := 10
	b := 3

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
}

func logicOperation() {
	a := true
	b := false

	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!(a && b))

}

func bitOperation() {

	fmt.Println(0 & 0)
	fmt.Println(0 | 0)
	fmt.Println(0 ^ 0)

	fmt.Println(0 & 1)
	fmt.Println(0 | 1)
	fmt.Println(0 ^ 1)

	fmt.Println(1 & 1)
	fmt.Println(1 | 1)
	fmt.Println(1 ^ 1)

	fmt.Println(1 & 0)
	fmt.Println(1 | 0)
	fmt.Println(1 ^ 0)

}

func valueCal() {
	a, b := 1, 2
	var c int
	c = a + b
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c = a + b, c =", c)

	plusAssignment(c, a)
	subAssignment(c, a)
	mulAssignment(c, a)
	divAssignment(c, a)
	modAssignment(c, a)
	leftMoveAssignment(c, a)
	rightMoveAssignment(c, a)
	andAssignment(c, a)
	orAssignment(c, a)
	norAssignment(c, a)
}

func plusAssignment(c, a int) {
	c += a // c = c + a
	fmt.Println("c += a, c =", c)
}

func subAssignment(c, a int) {
	c -= a // c = c - a
	fmt.Println("c -= a, c =", c)
}

func mulAssignment(c, a int) {
	c *= a // c = c * a
	fmt.Println("c *= a, c =", c)
}

func divAssignment(c, a int) {
	c /= a // c = c / a
	fmt.Println("c /= a, c =", c)
}

func modAssignment(c, a int) {
	c %= a // c = c % a
	fmt.Println("c %= a, c =", c)
}

func leftMoveAssignment(c, a int) {
	c <<= a // c = c << a
	fmt.Println("c <<= a, c =", c)
}

func rightMoveAssignment(c, a int) {
	c >>= a // c = c >> a
	fmt.Println("c >>= a, c =", c)
}

func andAssignment(c, a int) {
	c &= a // c = c & a
	fmt.Println("c &= a, c =", c)
}

func orAssignment(c, a int) {
	c |= a // c = c | a
	fmt.Println("c |= a, c =", c)
}

func norAssignment(c, a int) {
	c ^= a // c = c ^ a
	fmt.Println("c ^= a, c =", c)
}

func otherOperation() {

	a := 4
	fmt.Println("a=", a)
	var ptr *int
	ptr = &a
	fmt.Println("Address of a:", ptr)
	fmt.Println("Value of a through ptr:", *ptr)

}

func main() {
	fmt.Println("main method invoked!")
	intcal()
	changeTypeCal()
	calOperation()
	logicOperation()
	bitOperation()
	valueCal()
	otherOperation()
}

func intcal() {
	a, b := 1, 2
	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b

	fmt.Println("sum:", sum)
	fmt.Println("sub:", sub)
	fmt.Println("mul:", mul)
	fmt.Println("div:", div)

	c := 1
	c++
	fmt.Println("c after c++:", c)
	c--
	fmt.Println("c after c--:", c)

}

func changeTypeCal() {
	a := 3 + 0.9
	b := byte(1) + 1
	fmt.Println(a, b)

	sum := a + float64(b)
	fmt.Println(sum)

	sub := byte(a) - b
	fmt.Println(sub)

	mul := a * float64(b)
	fmt.Println(mul)

	div := int(a) / int(b)
	fmt.Println(div)
}
