package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	user := NewUser("Mike", 18)
	fmt.Println(user)
}
