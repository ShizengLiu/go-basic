package main

import (
	"fmt"
)

func main() {
	fmt.Println("main method invoked!")
	intVar()
	floatVar()
	stringVar()
	boolVar()
	zeroValue()
}

func intVar() {
	// 十六进制、八进制、二进制和十进制表示同一个数值
	var a uint8 = 0xF
	var b uint8 = 0xf

	var c uint8 = 017
	var d uint8 = 0o17
	var e uint8 = 0o17

	var f uint8 = 0b1111
	var g uint8 = 0b1111

	var h uint8 = 15

	fmt.Println(a, b, c, d, e, f, g, h)
	fmt.Println("a equals b:", a == b)
	fmt.Println("b equals c:", b == c)
	fmt.Println("c equals d:", c == d)
	fmt.Println("d equals e:", d == e)
	fmt.Println("e equals f:", e == f)
	fmt.Println("f equals g:", f == g)
	fmt.Println("g equals h:", g == h)
}

func floatVar() {
	var f1 float32 = 10
	f2 := 10.0
	fmt.Println(f1, f2)
	fmt.Println("f1 equals f2:", f1 == float32(f2))

	var z complex64 = 1.1 + 0.1i
	z2 := 1.1 + 0.1i
	z3 := complex(1.1, 0.1)
	fmt.Println("z:", z)
	fmt.Println("z2:", z2)
	fmt.Println("c3:", z3)
	fmt.Println("z equals z2:", z == complex64(z2))
	fmt.Println("z equals z3:", z == complex64(z3))
	x := real(z)
	y := imag(z)
	fmt.Println("real part:", x)
	fmt.Println("imaginary part:", y)

}

func stringVar() {
	var s string = "hello"
	var bytes []byte = []byte(s)
	fmt.Println("bytes:", bytes)

	var bytes1 []byte = []byte{97, 98, 99, 100}
	var s1 string = string(bytes1)
	fmt.Println("string:", s1)

	var r rune = '中'
	var r1 rune = 'a'
	fmt.Println("rune r:", r)
	fmt.Println("rune r1:", r1)

	var s2 string = "hello 中"
	var runes []rune = []rune(s2)
	fmt.Println("runes:", runes)

	var s3 string = "hello\nworld\n"
	var s4 string = `hello
world
`
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)
	fmt.Println("s3 equals s4:", s3 == s4)

	var s5 string = "Go语言"
	var bytes_s1 []byte = []byte(s5)
	var runes_s1 []rune = []rune(s5)
	fmt.Println("s5 length:", len(s5))
	fmt.Println("s5 bytes length:", len(bytes_s1))
	fmt.Println("s5 runes length:", len(runes_s1))

	fmt.Println("s5 sub:", s5[0:7])
	fmt.Println("s5 bytes sub:", string(bytes_s1[0:7]))
	fmt.Println("s5 runes sub:", string(runes_s1[0:3]))

}

func boolVar() {
	var b1 bool = true
	var b2 bool = false
	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b1 AND b2:", b1 && b2)
	fmt.Println("b1 OR b2:", b1 || b2)
	fmt.Println("NOT b1:", !b1)
	fmt.Println("NOT b2:", !b2)
}

func zeroValue() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("int zero value: %d\n", i)
	fmt.Printf("float64 zero value: %f\n", f)
	fmt.Printf("bool zero value: %t\n", b)
	fmt.Printf("string zero value: '%s'\n", s)
}
