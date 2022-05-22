# Commands

## To run a go program

`go run <filename>`

## To test

`go test`

## to watch

- install gow https://github.com/mitranim/gow#installation

# Syntax

Package declaration

```go
package main
```

Imports

```go
import "fmt"
```

Constants

```go
const spanishPrefix = "Hola, "
```

Global Variables

```go
var test = "Test"
```

Function creation + concatenation of string + return

```go
func Hello(name string) string {
  return "Hello, " + name
}
```

System Print

```go
fmt.Println("Hello, world")
```

If condition

```go
	if language == "" {
		language = "english"
	}
```

Initialize and assign variable

```go
x := 2
```

Overwrite, assign variable

```go
x = 2
```

Maps

```go
var prefixes = map[string]string{
	"spanish": "Hola, ",
	"french":  "Bonjour, ",
	"english": "Hello, ",
}

// use the map
language := "english"

prefixes["spanish"]
prefixes[language]
```

Test Format

```go
package main

import "testing"

func TestHello(t *testing.T) {
	// to not repeat got != want
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper() // this tells that when a test fails, the failure won't be in here but rather in the caller of this helper
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Harman", "")
		want := "Hello, Harman"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

```
