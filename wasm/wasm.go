package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	text := "Hello, WebAssembly World!!"

	js.Global().
		Get("document").
		Call("getElementById", "text").
		Set("innerText", text)

	fmt.Println(text)
}
