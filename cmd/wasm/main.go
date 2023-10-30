package main

import (
	"fmt"
	"syscall/js"
)

var htmlString = "<h2>Nothing like some html in a string</h2>"

func GetHtml() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		return htmlString
	})
}

func main() {
	ch := make(chan bool)
	fmt.Printf("call console.log from golang\n")
	js.Global().Set("getHtml", GetHtml())
	<-ch // block indefinitely to keep program alive
}
