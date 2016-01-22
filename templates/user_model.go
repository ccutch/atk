package TESTMODEL

import "github.com/ccutch/atk/core"

// User defined Model User
type User struct {
	core.Model

	Age  int    `json:"age"`
	Name string `json:"name"`
}

// Create pointer to a new instance of Model User
func NewUser(name string, age int) *User {
	m := make(User)

	m.Age = age
	m.Name = name
	return m
}
