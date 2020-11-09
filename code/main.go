package main

import (
	"fmt"
)

const (
	pi     = 3.1415
	first  = 1
	second = iota
	third  = iota
	fourth = 4 << iota
	fifth  = iota + 5
)

//availbale to the whole package
//have constant blocks!

func main() {
	//how to declare variables
	var i int
	i = 42

	// ^^ is verbose ~~~
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)
	// also verbose

	firstName := "Zeke"
	fmt.Println(firstName)
	//most of the time you want to do this

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)
	//complex data types!

	r, im := real(c), imag(c)
	//multiple assignments like python
	fmt.Println(r, im)

	var lastName *string = new(string) //init pointer
	*lastName = "Kei"
	fmt.Println(lastName)
	fmt.Println(*lastName)
	//pointer data type!
	//dereference with an asterix

	dog := "Woof"
	fmt.Println(dog)

	ptr := &dog
	fmt.Println(ptr, *ptr)

	dog = "cat"
	fmt.Println(ptr, *ptr)

	//address of is &

	/*
		iota and constant expressions
	*/
	fmt.Println(pi)
	fmt.Println(first, second, third, fourth, fifth)
	//iota is used? the increments by one
	// you don't have to explicitly called iota all of the time!
	//iota resets in iota blocks

	var arr [3]int
	//here's an array ~ size data type
	//simplify compiler
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	//kinda verbose tho
	cats := [3]int{1, 2, 3}
	fmt.Println(cats)

	//slice. array but flexible!

	slice := cats[:]
	//slice
	fmt.Println(slice)
	//slice is pointing to the data in the array

	pizza := []int{1, 2, 3}
	//compiler will know the size for the underlying array
	fmt.Println(pizza)

	pizza = append(pizza, 4, 42, 27)

	fmt.Println(pizza)
	//underlying array is handled by go

	s2 := pizza[1:]
	s3 := pizza[:2]
	//slices of slices
	fmt.Println(s2)
	fmt.Println(s3)
	s4 := pizza[1:2]
	fmt.Println(s4)
}
