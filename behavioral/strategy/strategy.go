package main

import "fmt"

// Strategy 策略
type Strategy interface {
	Apply(int, int) int
}

// Operation 包装类
type Operation struct {
	Strategy Strategy
}

func NewOperation(strategy Strategy) *Operation {
	return &Operation{
		Strategy: strategy,
	}
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Strategy.Apply(leftValue, rightValue)
}

type Add struct {
}

func (a Add) Apply(l, r int) int {
	return l + r
}

var _ Strategy = new(Add)

type Multiplication struct{}

func (Multiplication) Apply(l, r int) int {
	return l * r
}

var _ Strategy = new(Multiplication)

func main() {
	add := Add{}

	fmt.Println(NewOperation(add).Operate(1, 2))

	multi := Multiplication{}

	fmt.Println(NewOperation(multi).Operate(1, 2))
}
