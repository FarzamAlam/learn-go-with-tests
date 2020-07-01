package main

import "fmt"

func main() {
	fmt.Println(Hello("Dina"))
}

const helloPrefix = "Hello"

func Hello(name string) string {
	if name == "" {
		name = "World!"
	}
	return helloPrefix + ", " + name
}
