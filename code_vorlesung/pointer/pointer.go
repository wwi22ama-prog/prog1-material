package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (person *Person) SetAge(age int) {
	if age > 0 && age < 120 {
		person.Age = age
	}
}

func main() {
	var zahl1 int = 41
	Add1(zahl1)
	fmt.Println(zahl1)

	Add1_Pointer(&zahl1)
	fmt.Println(zahl1)

	p1 := Person{"Max Mustermann", 42}
	p1.SetAge(99)

	fmt.Println(p1)
}

func Add1(x int) {
	x++
}

func Add1_Pointer(x *int) {
	(*x)++
}
