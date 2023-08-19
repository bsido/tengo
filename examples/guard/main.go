package main

import (
	"fmt"

	"github.com/bsido/tengo"
)

func main() {
	src := `
	control := 0

	function := func(fail) {
		if fail {
			return error("smth failed")
		} else {
			return "ok"
		}
	}
 	// this is fine
    guard ok := function(false)

	// this will exit with an error
	guard nok := function(true)

	// we will never get here
	control = 1
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
