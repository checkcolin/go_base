package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description    interfaces and functions
 **/

type (
	ID     uint64
	SSN    string
	Person struct {
		name string
		age  uint8
		ssn  SSN
	}
	Empty interface{}
)
func main() {
	// review - declaring interface variables and assigning values
	var e Empty
	fmt.Printf("e's value: %v, type: %T\n", e, e)
	e = 11.04
	fmt.Printf("e's value: %v, type: %T\n", e, e)
	e = ID(19721104)
	fmt.Printf("e's value: %v, type: %T\n", e, e)
	e = SSN("019-72-1104")
	fmt.Printf("e's value: %v, type: %T\n", e, e)
	e = &Person{name: "Jane Doe", age: 35, ssn: SSN("019-72-1104")}
	fmt.Printf("e's value: %v, type: %T\n", e, e)
}