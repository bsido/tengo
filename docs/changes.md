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

## The `exit` statement

The `exit` statement is used to exit the script either with or without an error.
This statement can be called with zero or one argument.

```
exit
```

which is equivalent to the following `tengo` code:

```go
os := import("os")

os.exit(0)
```

> Note: the above function `os.exit` was removed from the standard library in this fork.

If the user wants to exit with an error, the `exit` statement can be called with one argument:

```
exit <expression>
```

The `expression` must resolve to a `string` or an `error` value, otherwise the script will fail with an unexpected runtime error.
Examples:

```go
exit "error message"

exit error("error message")

function := func() {
    return error("error message")
}

exit function()
```

In the above cases the script will exit with an exit code `1` and print the error message, somewhat ~equivalent to:

```go
os := import("os")
fmt := import("fmt")

os.println("error message")
os.exit(1)
```
