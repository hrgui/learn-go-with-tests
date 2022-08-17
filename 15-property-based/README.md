# strings.Builder

A builder is used to efficiently build a string using Write methods. It minimizes memory copying.

```go
package main

import "strings"

func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for i := 0; i < arabic; i++ {
		result.WriteString("I")
	}

	return result.String()
}
```

# a for loop like a while

```go
func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for arabic > 0 {
		switch {
		case arabic > 4:
			result.WriteString("V")
			arabic -= 5
		case arabic > 3:
			result.WriteString("IV")
			arabic -= 4
		default:
			result.WriteString("I")
			arabic--
		}
	}

	return result.String()
}

```

# An intro to property based tests

```go
func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic int) bool {
   if arabic < 0 || arabic > 3999 {
   	log.Println(arabic)
   	return true
   }
   roman := ConvertToRoman(arabic)
   fromRoman := ConvertToArabic(roman)
   return fromRoman == arabic
}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
```

Property based tests help take rules we know about or domain and somehow exercise them against our code.

In go, we can use `testing/quick` package from the standard library.

We use `quick.Check`:

> Check looks for an input to f, any function that returns bool, such that f returns false. It calls f repeatedly, with arbitrary values for each argument. If f returns false on a given input, Check returns that input as a \*CheckError. For example:

```go
func TestOddMultipleOfThree(t *testing.T) {
	f := func(x int) bool {
		y := OddMultipleOfThree(x)
		return y%2 == 1 && y%3 == 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
```

- built into standard library
- can give more confidence
- force to think about the domain deeply
- nice complement to test suite
