package main

import "fmt"

func main() {
	fmt.Println(Hello("Dina", ""))
}

const english = "Hello"
const spanish = "Hola"
const french = "Bonjour"

func Hello(name string, language string) string {
	if name == "" {
		name = "World!"
	}

	return greetingPrefix(language) + ", " + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "french":
		prefix = french
	case "spanish":
		prefix = spanish
	default:
		prefix = english
	}
	return
}
