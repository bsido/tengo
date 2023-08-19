package main

import (
	"fmt"

	"github.com/bsido/tengo"
)

func main() {
	exitWithoutArgument()

	exitWithMessage()

	exitWithError()

	exitInsideFunction()
}

func exitInsideFunction() {
	src := `
control := 0

function := func(fail) {
	if fail {
		exit "smth failed"
	} else {
		return "ok"
	}
}

ok := function(false)
nok := function(true)
`

	script := tengo.NewScript([]byte(src))
	compiled, err := script.Run()

	control := compiled.Get("control")
	ok := compiled.Get("ok")
	nok := compiled.Get("nok")

	fmt.Println("control: ", control)
	fmt.Println("ok: ", ok)
	fmt.Println("nok: ", nok)
	fmt.Println("err: ", err)
}

func exitWithoutArgument() {
	src := `
control := 0

exit

// we will never get here
control = 1
`

	script := tengo.NewScript([]byte(src))
	compiled, err := script.Run()

	control := compiled.Get("control")

	fmt.Println("control: ", control)
	fmt.Println("err: ", err)
}

func exitWithMessage() {
	src := `
control := 0

exit "smth failed"

// we will never get here
control = 1
`

	script := tengo.NewScript([]byte(src))
	compiled, err := script.Run()

	control := compiled.Get("control")

	fmt.Println("control: ", control)
	fmt.Println("err: ", err)
}

func exitWithError() {
	src := `
control := 0

exit error("exited with error")

// we will never get here
control = 1
`

	script := tengo.NewScript([]byte(src))
	compiled, err := script.Run()

	control := compiled.Get("control")

	fmt.Println("control: ", control)
	fmt.Println("err: ", err)
}
