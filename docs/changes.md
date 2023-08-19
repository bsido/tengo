# Changes to the original code

This document contains all changes that were made to the language after forking the original repository.

## The `quard` statement

The `guard` statement is a conditional statement that makes the script exit with a `1` exit code if the condition is not met.
The `quard` statement is used when assigning a value to a variable.

```go
quard x := expression
```

Even though the `expression` can be anything, the `guard` statement is only useful when it is a function returning some value or an error.
The `guard` statement is equivalent to the following `tengo` code:

```go
"os" := import("os")
"fmt" := import("fmt")

x := call_function()
if is_error(x) {
    fmt.println(x.value)
    os.Exit(1)
}
```

and to this `go` code:

```go
if err := call_function(); err != nil {
    println(err)
	os.Exit(1)
}
```
