package test_module

import "fmt"

func init() {
	fmt.Println("You are using test_module")
}

func Hey() {
	fmt.Println("Jude")
}

func Hoy(data int) {
	fmt.Println("Not ", data)
}
