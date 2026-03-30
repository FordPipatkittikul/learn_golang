package main

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

func main() {

	c := CustomerRepositroyMock{}
	c.On("GetCustomer", 1).Return("Ford", 18, nil)
	c.On("GetCustomer", 2).Return("", 0, errors.New("not found"))

	name, age, err := c.GetCustomer(1)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(name, age)
}

type CustomerRepositroy interface {
	GetCustomer(id int) (name string, age int, err error)
	Hello()
}

type CustomerRepositroyMock struct {
	mock.Mock
}

func (m *CustomerRepositroyMock) GetCustomer(id int) (name string, age int, err error) {
	args := m.Called(id)
	return args.String(0), args.Int(1), args.Error(2)
}

func (m *CustomerRepositroyMock) Hello() {
	println(m)
}