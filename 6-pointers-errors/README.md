# Pointers and errors

## Pointers

In Go, it copies values when you pass them to functions and methods. If you are writing a function that needs to mutate state, change the argument to a pointer to the thing you want to change. Otherwise all changes only reflect the copy.

## nil

- Pointers can be nil
- When a function returns a pointer to something, you need to make sure if it's nil or you might raise a runtime exception.
- `null` in JavaScript

## Errors

- Errors are the way to signify failure when calling a function or method.
- https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
