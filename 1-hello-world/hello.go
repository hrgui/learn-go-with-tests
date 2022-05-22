package main

import (
	"fmt"
	"strings"
)

var prefixes = map[string]string{
	"spanish": "Hola, ",
	"french": "Bonjour, ",
	"english": "Hello, ",
}


func Hello(name string, language string) string {
	language = strings.ToLower(language)

	if name == "" {
		name = "World"
	}

	if language == "" {
		language = "english"
	}

	return prefixes[language] + name
}

func main() {
	fmt.Println(Hello("Harman", ""))
}
