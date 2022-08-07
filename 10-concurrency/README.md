# Concurrency

In Go, we can make operations concurrent by utilizing goroutines.

## Goroutines and anonymous functions

Basic unit of concurrency in Go. Its signature looks like the following:

```go
go func() {

}()
```

It utilizes an anonymous function - an unnamed function `func() {}`.

## Channels

Channels are a Go data structure that can both receive and send values. These operations, along with details allow communication between different processes.

Channels are declared like so

```go
resultChannel := make(chan result)
```

We use the `<-` operator to send a result struct to the resultChannel

```go
type result struct {
	string
	bool
}

// wc is WebsiteChecker in code, but assumed can be anything
resultChannel <- result{u, wc(u)}
```

We can also send the `<-` operator to receive:

```go
r := <-resultChannel
results[r.string] = r.bool
```

## Detecting Races

We can detect races in our tests by doing the following:

```
go test -race
```

## Running benchmarks

```
go test -bench=.
```

- suffix it with `benchmark_test.go`
- prefix each test with `Benchmark`
