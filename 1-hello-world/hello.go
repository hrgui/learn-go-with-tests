package main

import (
	"fmt"
	"strings"
)

const englishHelloPrefix = "Hello, "

const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	language = strings.ToLower(language)

	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case "french":
		prefix = frenchHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Harman", ""))
}
