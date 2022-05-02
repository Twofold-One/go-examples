package basic

import "fmt"

func ArraysExample() {
	// basic array creation
	var myArr [2]int
	myArr[0] = 1
	myArr[1] = 2
	fmt.Println(myArr)

	// array literal
	myArrL := [2]int{2, 3}
	fmt.Println(myArrL)

	// size computed by the compiler
	myArrS := [...]int{4, 5, 6}
	fmt.Println(myArrS)

	// length and capacity of array
	myArr1 := [...]string{"one", "two", "three"}
	fmt.Println(myArr1, len(myArr1), cap(myArr1))

	// accessing array elements
	a := [3]int{1, 2, 3}
	firstElement := a[0]
	secondElement := a[1]
	fmt.Println(firstElement, secondElement)

	// iterating over an array
	b := [5]string{"a", "b", "c", "d", "e"}
	for i, v := range b {
		fmt.Printf("value at index %d is %s\n", i, v)
	}

	// ingoring index value
	c := [3]int{1, 2, 3}
	for _, v := range c {
		fmt.Println(v)
	}

	d := [5]int{1, 2, 3, 4, 5}

	// ascending iteration
	for i := 0; i < len(d); i++ {
		fmt.Println(i, d[i])
	}

	// descending iteration
	for i := len(d) - 1; i >= 0; i-- {
		fmt.Println(i, d[i])
	}

	fmt.Println(getIndex(d, 3))

	// passing array to functions

	testArr := [2]string{"i1", "i2"}

	//without modification of initial array, just copying
	fmt.Println(testArr, UpdateArrayWithoutMod(testArr), testArr)
	// modifying initial array
	fmt.Println(testArr, UpdateArrayWithMod(&testArr), testArr)

	// multidimensional arrays
	firstRow := [3]int{1, 2, 3}
	secondRow := [3]int{4, 5, 6}
	thirdRow := [3]int{7, 8, 9}

	matrixArr := [3][3]int{firstRow, secondRow, thirdRow}
	fmt.Println(matrixArr)
}

// finding element inside an array
func getIndex(haystack [5]int, needle int) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

func UpdateArrayWithoutMod(arr [2]string) string {
	arr[0] = "updated"
	return "arr has been copied"
}

func UpdateArrayWithMod(arr *[2]string) string {
	arr[0] = "updated"
	return "arr modified"
}