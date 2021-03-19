package main // Declares the packages you're using

import (
	"errors"
	"fmt"
	"math"
) // Imports different packages

func main() {
	fmt.Println("Hello World") // Prints then makes a new line
	fmt.Print("Hello World")   // Prints and doesn't make a new line
	fmt.Println()

	var x int = 10 // Declaring variables, states its a variable then the name and the data type following
	var y int      // Variables can be declared without assingment
	z := 5         // Initialises variables like in python, Interprets the data type
	fmt.Println(y + x + z)

	// Array
	var p [2]int // Long hand declaration
	p[0] = 3     // Same way as python to reference values
	p[1] = 6
	fmt.Println(p)

	q := []int{3, 6, 5, 2}
	fmt.Println(q)
	fmt.Println(q[1])
	q = append(q, 3)

	// Maps
	verticies := make(map[string]int) // Basically a dictionary
	verticies["Triangle"] = 3
	verticies["Square"] = 4
	verticies["Circle"] = 0
	fmt.Println(verticies)
	fmt.Println(verticies["Square"]) // Prints the value associated to the key
	delete(verticies, "Square")      // Removes Square

	// Loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	} // Only loop in Go is the for loop

	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	} // However it can be converted into a while loop

	arr := []int{4, 3, 2}
	for index, value := range arr {
		fmt.Println("index: ", index, "value: ", value)
	}

	// Using Functions, you have to state the data type and state what value you're returning. You can return more than one type incase of an error
	r := sum(6, 8)
	fmt.Println(r)

	result, err := sqrt(16) // When you call this you can get two values hence the two variable names

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Using Structs
	s := person{name: "Ben", age: 17} // Need to use curly brackets here
	fmt.Println(s.age)

	// Pointers
	t := 5
	inc(&t)
	fmt.Println(t) // Gets the memory location

}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

func inc(x *int) {
	*x++
}

// Creating Structs: Basically a btec class
type person struct {
	name string
	age  int
}
