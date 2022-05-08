package basic

import (
	"fmt"
)

func EnumExample() {
	// enum or enumeration data type is data type that consists of a set
	// of values that are explicitly defined by the programmer

	// not explicitly supported in go v1
	// https://github.com/alvaroloes/enumer
	// https://github.com/abice/go-enum
}

type HTTPMethod int

const (
	// iota is an identifier which is used with constant and which can simplify constant definitions that use auto increment numbers
	// initial value of iota is zero
	GET HTTPMethod = iota
	POST
	PUT
	DELETE
	PATCH
	HEAD
	OPTIONS
	TRACE
	CONNECT
)

type TestEnum int

const (
	First TestEnum = iota + 1
	Second
	Third
)

type TestEnum2 int

const (
	One TestEnum2 = iota * 2
	Two
	Three
)

func IotaExample() {
	fmt.Println(PUT)
	fmt.Println(First, Second, Third)
	fmt.Println(One, Two, Three)
}

func BitwiseOperationsExample() {
	fmt.Printf("\nBitwise Operations Examples\n")
	// can be used only with integers
	// to manipulate bits we always start at right side
	var x, y, z uint8
	// AND & bitwise operator
	x = 1
	y = 2
	z = x & y
	fmt.Printf("%08b\n%d\n", z, z)

	// OR | bitwise operator
	x = 200
	y = 100
	z = 200 | 100
	fmt.Printf("%08b\n%d\n", z, z)

	// XOR (Xclusive OR) ^ bitwise operator
	// x ^ y is equal to 1 only if one of the operands is equal to 1, not both.x ^ y is equal to 1 only if one of the operands is equal to 1, not both\
	x = 200
	y = 100
	z = 200 ^ 100
	fmt.Printf("%08b\n%d\n", z, z)

	//  NOT ^ bitwise operator
	x = 200
	z = ^x
	fmt.Printf("%08b\n%08b\n%d\n", x, z, z)

	// AND NOT &^ bitwise operator
	// x AND NOT y / x AND (NOT y)

	// number of positions
	var n int

	// left shift
	x = 200
	fmt.Printf("%08b\n", x)
	n = 1
	z = x << n
	fmt.Printf("%08b\n%d\n", z, z)

	// we are left-shifting the byte 11001000 by one position. The result of this shift is 10010000. We have added a zero at the left of the byte and shifted the other by 1 position
	// we are storing our integers into 8 bits, so we are losing bits. to avoid that, we can store our numbers on 16 bits (2 bytes)
	var u, v, w uint16
	u = 200
	fmt.Printf("%08b\n", u)
	v = 1
	w = u << v
	fmt.Printf("%08b\n%d\n", w, w)
	v = 2
	w = u << v
	fmt.Printf("%08b\n%d\n", w, w)
	// when you left shift the binary representation of a number by n position you multiply it’s decimal representation bt it by two at the power n

	// right shift
	u = 200
	fmt.Printf("%08b\n", u)
	v = 3
	w = u >> v
	fmt.Printf("%08b\n%d\n", w, w)
	// you can see here that we are dividing 200 by 3 when we shift the bytes to the right. This is another property of binary numbers. When you left shift the binary representation of a number (in base 10) by n position, you divide it by two at the power n
}