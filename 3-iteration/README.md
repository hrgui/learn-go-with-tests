Iteration

```go
for i := 0; i < b.N ; i++ {

}
```

Run the benchmark

```go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		Repeat("a")
	}
}

```
