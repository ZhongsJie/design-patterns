package main

import "fmt"

type Speed string

const (
	MPH Speed = "1"
	KPH       = "1.60934"
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

type car struct {
	color  Color
	wheels Wheels
	speed  Speed
}

func (c car) String() string {
	return string(c.color) + "/" + string(c.wheels) + "/" + string(c.speed)
}

func (c car) Drive() error {
	fmt.Printf("car %s: Drive", c)
	return nil
}

func (c car) Stop() error {
	fmt.Printf("car %s: Stop", c)
	return nil
}

func (c car) Color(color Color) Builder {
	c.color = color
	return c
}

func (c car) Wheels(wheels Wheels) Builder {
	c.wheels = wheels
	return c
}

func (c car) TopSpeed(speed Speed) Builder {
	c.speed = speed
	return c
}

func (c car) Build() Interface {
	return c
}

var _ Interface = new(car)
var _ Builder = new(car)

func NewBuilder() Builder {
	return &car{}
}

func main() {
	car := NewBuilder().Color(RedColor).TopSpeed(KPH).Wheels(SportsWheels).Build()

	_ = car.Drive()

	_ = car.Stop()
}
