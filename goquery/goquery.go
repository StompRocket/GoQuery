package goquery

import (
	"syscall/js"
	"errors"
	"strings"
)

// This type contains a DOM object, and Go abstractions of it's 
// various properties and values. Methods can be called on these
// to ease development. This way you don't have to do all the
// syscall/js nonsence manually.
type DomObject struct {
	object js.Value
	ClassList ClassList
	InnerHTML string
}

// Find an element that matches the query parameter, and return a
// GoQuery object containing that DOM object and it's relevant info. 
// Additional opperations can be made on the DomObject type after it
// is created with the Q function.
//
// It's called Q because $ is not an allowed function name and It is
// still pretty short. 
func Q(selector string) (DomObject, error) {
	el := js.Global().Get("document").Call("querySelector", selector)
	if el.Type() == js.TypeObject {
		cl := el.Get("classList")
		obj := DomObject{
			object: el,
			ClassList: ClassList{
				object: cl,
				Length: cl.Get("length").Int(),
				Value: cl.Get("value").String(),
				Items: strings.Split(cl.Get("value").String(), " "),
			},
			InnerHTML: el.Get("innerHTML").String(),
		}
		return obj, nil
	}
	return DomObject{}, errors.New("Failed to get DOM object")
}
