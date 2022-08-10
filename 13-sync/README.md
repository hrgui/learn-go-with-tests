# Sync

When running functions concurrently that mutate state, we have to think - ok if everyone runs the code at the same time, what happens?

State won't be managed correctly and won't be what's expected.

It is best to use a mutex!

```go

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}
```

- Use channels when passing ownership of data.
- Use mutexes for managing state.

# go vet

It's like eslint but for go. It alerts to some subtle bugs in your code before they hit poor users.
