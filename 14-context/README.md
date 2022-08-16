# Contexts

Also followed this youtube video: https://www.youtube.com/watch?v=h2RdcrMLQAo

Context can be used for the following

- to store key/value pairs
- to manage processes

For managing processes, we use `ctx.Done()` and `cancel()`. It's best for timing out something.

- ctx.WithCancel / ctx.WithTimeout gives the idea of timing out

Use the storage for key/value pairs only for trace ids; values that are orthogonal to a request. `ctx.Value()` should only inform, not control. It should never be required input for documented or expected results.

References

- https://go.dev/blog/context (try out the google web search example)
