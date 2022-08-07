# select

- Allows you to wait on multiple channels.
- synchronizes processes

Remember that we can wait for values to be sent to a channel with `myVar := <-ch`. However, this is a blocking call, as we're waiting for a value.

In a select, the first one to send a value "wins" and code underneath the case is executed.

```go
func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
```

# time.After

Returns a channel, and will send a signal down it after the amount of time you define.

```go
func Racer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}
```

In the above, if we've timedout, time.After wins and the case will be invoked.

# httptest

- It is a convenient way of creating test servers so you can have reliable and controllable tests.
- Uses the same interfaces such as the "real" net/http servers, which is consistent and less for you to learn.
