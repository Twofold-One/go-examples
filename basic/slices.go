package basic

import (
	"fmt"
	"strings"
)

func SlicesExample() {
	// creating empty slice with len = 2
	s := make([]int, 2)
	s[0] = 1
	s[1] = 2
	fmt.Println(s)

	// slice literal
	s1 := []int{1, 2, 3}
	fmt.Println(s1)

	// slicing an array
	numArr := [4]string{"one", "two", "three", "four"}
	// slicing including element 0 to element 2 (exluding or element 2 - 1)
	numS1 := numArr[0:2]
	fmt.Println(numS1)

	numS2 := numArr[2:]
	fmt.Println(numS2)

	// slicing does not copy data
	testArr1 := [3]string{"a", "b", "c"}
	testSlice1 := testArr1[0:1]
	fmt.Println(testSlice1)
	// modify original array
	testArr1[0] = "b"
	fmt.Println("After modification of original array")
	fmt.Println(testSlice1)

	// slicing a string
	room := "bedroom"
	roomS := room[3:]
	fmt.Println(roomS)

	// slice len (uint) and cap (uint) (inherits from original array)
	a1 := [5]int{1, 2, 3, 4, 5}
	a1S := a1[0:3]
	fmt.Printf("length of a1S is %d and capacity is %d\n", len(a1S), cap(a1S))

	// slices in function parameters
	// a function that takes a slice as parameter can change the underlying array
	// because a slice is internally a pointer to an underlying array
	s2 := []int{1, 2, 3}
	multiply(s2, 2)
	fmt.Println(s2)

	// if appending to a slice in the function
	// it's better to use pointer to it
	s3 := []string{"start", "mid"}
	addEnd(&s3)
	fmt.Println(s3)

	// make func to create slice with certain length and capacity
	s4 := make([]int, 2, 4)
	fmt.Println(len(s4), cap(s4))

	// copy func: copy (dst, src []Type) int
	// destination > source
	a := []int{10, 20, 30, 40}
	b := []int{1, 1, 1, 1, 1}
	copy(b, a)
	fmt.Printf("a: %v, b: %v\n", a, b)
	// destination < source
	a = []int{10, 20, 30 , 40}
	b = []int{1, 1}
	copy(b, a)
	fmt.Printf("a: %v, b: %v\n", a, b)
	// destination = source will copy as it is

	// append func: append(slice []Type, elems ...Type) []Type
	c := []int{1, 2, 3, 4}
	c = append(c, 5)
	fmt.Println(c)

	// for loops that modify a slice
	names := []string{"John", "Bob", "Gregory", "Evan"}
	for i := range names {
		names[i] = strings.ToUpper(names[i])
	}
	fmt.Println(names)

	// merge two slices
	sl1 := []int{1, 2, 3}
	sl2 := []int{4, 5}
	merge := append(sl1, sl2...)
	fmt.Println(merge)

	// remove an element at index
	// a = append(a[:i], a[i+1:]...)
	// example remove repeated number 3
	sl3 := []int{1, 2, 3, 3, 4, 5}
	sl3 = append(sl3[:2], sl3[3:]...)
	fmt.Println(sl3)

	// put an element at index
	// example insert 22 as index 1 element in slice
	sl4 := []int{1, 2, 3, 4, 5}
	// add element to the end of the slice
	sl4 = append(sl4, 0)
	// copy s[i:] to s[i+1:]
	fmt.Printf("dst: %v, src: %v\n", sl4[2:], sl4[1:])
	copy(sl4[2:], sl4[1:])
	fmt.Println(sl4)
	// set 22 at index 1
	sl4[1] = 22
	fmt.Println(sl4)

	// remove all elents from slice
	sl5 := []int{1, 2, 3}
	sl5 = sl5[:0]
	fmt.Println(sl5)

	// set slice to nil
	sl6 := []int{1, 2, 3, 4, 5}
	sl6 = nil
	fmt.Println(sl6)

	// prepend element to slice
	sl7 := []int{2, 3}
	sl7 = append([]int{1}, sl7...)
	fmt.Println(sl7)

	// dynamic creation of two-dimensional slices
	my2DSlice := [][]int{}
	for i := 0; i < 10; i++ {
		my2DSlice = append(my2DSlice, []int{i})
	}
	fmt.Println(my2DSlice)
}

func multiply(s []int, factor int) {
	for i := 0; i < len(s); i ++ {
		s[i] = s[i] * factor
	}
}

func addEnd(s *[]string) {
	*s = append(*s, "end")
}