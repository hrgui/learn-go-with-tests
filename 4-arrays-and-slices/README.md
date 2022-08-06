# Difference between array and slices

- **array** allows you to store multiple elements of the same type in a variable in a particular order. It has a **fixed** size.
- **slice** are the same as arrays except it has a variable size.

# useful array functions

- len(array) - length of the array
- append(array, x) - append a value to the end of the array
- copy() - copies the array
- x[a:b] - _slice_ the array x from a,b - but unlike in JS this is a pointer. Use copy to make a new one.

- make(x, length) - makes a slice with length

# Initialize an array

- they're fixed, so in example below n = 5

```go
numbers := [5]int{1, 2, 3, 4, 5}
```

Iteration over an fixed array

```go
func Sum(numbers [5]int) int {
	sum := 0
	for i :=0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}

```

Using range (similar to for of in JS)

```go
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

```

`for index, number := range numbers`
`_` is blank

# A slice

Omit the size and it's a slice. However, it doesn't mean the array is infinite. It's defined as how many values are defined in the slice.

```go
numbers := []int{1, 2, 3, 4, 5}
```

# Do coverage

```
go test -cover
```

# Varadic function arguments

- `SumAll(a,b)`
- `SumAll(a,b,c)`

```go
func SumAll(numbersToSum ...[]int) []int {
	sums := []int{}

	for _, numbers := range numbersToSum {
		sums = append(sums,Sum(numbers))
	}

	return sums
}
```
