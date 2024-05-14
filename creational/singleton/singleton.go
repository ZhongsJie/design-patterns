package main

import (
	"fmt"
	"sync"
)

type Singleton map[string]string

var (
	once     sync.Once
	instance Singleton
)

func New() Singleton {
	once.Do(func() {
		instance = make(Singleton)
	})
	return instance
}

func main() {
	singleton := New()
	singleton["name"] = "xiaoming"

	singleton2 := New()
	fmt.Print(singleton2["name"])

}
