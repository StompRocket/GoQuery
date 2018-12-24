package goquery

import (
	"syscall/js"
	"errors"
	"strings"
)

type DomObject struct {
	object js.Value
	ClassList ClassList
	InnerHTML string
}

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
