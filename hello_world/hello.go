package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	english = "English"
	french  = "French"

	spanishHelloPrefix = "Hola, "
	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
)

func greetingsPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingsPrefix(language) + name

}

func main() {
	fmt.Println(Hello("John", "English"))
}
