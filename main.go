package main

import (
	"fmt"

	"github.com/elct9620/demo-stdio-go-plugin/internal/plugin"
)

func main() {
	datePlugin, err := plugin.NewPlugin("/bin/date")
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := datePlugin.Call("", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	outStr, ok := out.([]byte)
	if !ok {
		fmt.Println("Error: output is not a byte slice")
		return
	}

	fmt.Print(string(outStr))
}
