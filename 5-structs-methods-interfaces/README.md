# Structs, Methods and Interfaces

# Function

Recall functions are a block of organized, reusable code that is used to perform a single, related action. It may accept 0 to many arguments, and it may return a output value.

An example of a function is the following:

```go
func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}
```

An argument is written in this format:

```go
width float64
```

# Structs

A struct is a named collection of fields where you can store data. In JavaScript, this is called an object.

```go
type Rectangle struct {
	Width  float64
	Height float64
}
```

Like functions, structs can also be declared inline, or anonymous:

```go
areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}
```

# Methods

A method is a function of a struct. An example of a method is the following:

```go
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
```

Notice that we begin with the struct in question, then the method name. This associates the method to the struct.

We can then call the method like so:

```go
rectangle := Rectangle{Width, 12, Height: 6}
rectangle.Area()
```

# Interfaces

Interfaces are like blueprints of a struct. They hold what properties a struct could possibly have, but it may not contain all properties.

Interfaces allow for reusability of code when we have multiple structs that have similar methods, and we are only concerned about the similar methods.

See the following example:

```go
package main

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}
}

```

In this example:

- We declared an anonymous struct array of `{shape Shape, want float64}`for our tests.
- We test the `Area()` function of the interface. Notice we are not concerned that a Circle has a radius, a Rectangle has a width and a height, and a Triangle has a base and a height.
