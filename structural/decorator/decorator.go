package main

import "log"

type HandleFunc func(int) int

func LogDecorate(fn HandleFunc) HandleFunc {
	return func(n int) int {
		log.Println("Starting the execution with the integer", n)

		result := fn(n)

		log.Println("Execution is completed with the result", result)

		return result
	}
}

func main() {
	handler := func(in int) int {
		return in
	}
	LogDecorate(handler)(1)
}
