package dom

import (
	"syscall/js"
)

// DOM represents a document object model that will be rendered in the browser
type DOM struct {
	Header any
	Body   any
}

// ok so first implementation should just be silly
// how do I define a div?
var htmlString = "<h2>here is my new site</h2>"

func GetHtml() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		return htmlString
	})
}

// so there has to be a render method
func (d *DOM) Render() {
	js.Global().Set("getHtml", GetHtml())
}
