package users

import (
	"fmt"
	"math/rand"
)

//embeber struct

type User struct {
	Id   int
	Name string
}

type Employee struct {
	User
	Active bool
}

type Cashier interface {
	CalcTotal(item ...float32) float32
	desactivate()
	reactivate()
}

type Admin interface {
	DesactivateEmployee(c Cashier)
}

func NewEmploye(name string) *Employee {
	return &Employee{
		User: User{
			Id:   rand.Intn(1000),
			Name: name,
		},
		Active: true,
	}
}

func (e *Employee) CalcTotal(item ...float32) float32 {

	if !e.Active {
		fmt.Println("Cashier is desactivated")
		return 0
	}

	var sum float32

	for _, i := range item {
		sum += i
	}

	return sum + 1.15
}

func (e *Employee) desactivate() {
	e.Active = false
}

func (e *Employee) DesactivateEmployee(c Cashier) {
	c.desactivate()
}

func (e *Employee) reactivate() {
	e.Active = true
}
