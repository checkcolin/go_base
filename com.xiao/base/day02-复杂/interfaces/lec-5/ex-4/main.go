package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description
 **/

type (
	Person struct {
		name string
	}
)

func main() {
	var foo interface{}
	foo = getValue()
	//类型判断
	fmt.Printf("foo's value: %v, type: %T\n", foo, foo)
	intVal,ok:=foo.(int)
	if ok{
		fmt.Printf("intVal's is type 'int' with intValue: %v\n", intVal)
	}
	floatVal, ok := foo.(float64)
	if ok{
		fmt.Printf("floatVal's is type 'float64' with floatValue: %v\n", floatVal)
	}

	strVal, ok := foo.(string)
	if ok{
		fmt.Printf("strVal's is type 'string' with strValue: %v\n", strVal)
	}

	pPerson, ok := foo.(*Person)
	if ok{
		fmt.Printf("pPerson's is type '*Person' with pPersonue: %v\n", pPerson)
	}

}
func getValue()interface{}{
	ch:=make(chan interface{},1)
	select {
	case ch<-2021:
	case ch<- 9.15:
	case ch<- "Hello world!":
	case ch <- &Person{name: "Jane Doe"}:
	}
	return <-ch
}
