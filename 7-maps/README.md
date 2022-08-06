# Maps

Maps allow you to store items via a key and look up them quickly.

# Example: How to declare a map

```go
dictionary := map[string]string{"test": "this is just a test"}
```

# How to lookup a map

```go
dictionary[word]
```

Note that if the key doesn't exist in the map, it will return nothing back, e.g. `""`

It returns a tuple back of `value, found`, where `value` is the type of the value and found is a boolean indicating if it was found or not.

```go
value, ok = dictionary[word]
```

# How to update or add a value to the map

```go
dictionary[word] = updatedValue
```

# How to delete the key,value entry

Go has a special function called `delete` which takes in the dictionary and the key to look for.

```go
delete(dict, key)
```

Note that if the key is not there, it has no effect.
