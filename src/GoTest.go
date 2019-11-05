package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age  int
}

func main() {
	var x int = 5
	y := 7
	sum := x + y
	fmt.Println(sum)

	// If
	if x > 6 {
		fmt.Println("More than 6")
	}

	// Arrays
	var a [5]int
	b := [5]int{5, 4, 3, 2, 1} // Can't add
	c := []int{6, 5, 7, 8, 9}  // Can use append to add
	a[2] = 7
	fmt.Println(a, b, c)

	// Mappings
	vertices := make(map[string]int)

	vertices["traignle"] = 2
	vertices["sqaure"] = 3
	vertices["circle"] = 12

	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])
	delete(vertices, "square")

	// Loops

	// For loop: Only loop in Go.
	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}

	// While loop using for loop.
	j := 0

	for j < 5 {
		fmt.Print(j)
		j++
	}

	// Using range with arrays
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index: ", index, "value: ", value)
	}

	//Using range with maps
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key: ", key, "value: ", value)
	}

	// Calling function: sum
	result := sum(2, 3)
	fmt.Println(result)

	// Calling function sqrt.
	result2, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Structs
	p := person{name: "Jake", age: 23}
	fmt.Println(p.age)

	// Pointers
	i := 7
	fmt.Println(i)  // Value of the Variable
	fmt.Println(&i) // Memory Address of the Varible

	fmt.Println(increment(i)) // This is useless since i is copied by value

	incRes := incrementWell(&i)

}

func sum(x, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefinded for negative numbers")
	}

	return math.Sqrt(x), nil
}

func increment(y int) {
	y++
}

func incrementWell(x *int) {
	*x++ //De-referencing (So that you won't increment the memory address, but what's inside the memory address) the pointer and incrementing the actual variable.
}
