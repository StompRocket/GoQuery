package goquery

import (
	"syscall/js"
)

type ClassList struct {
	object js.Value
	Length int
	Value string
	Items []string
}

func (c *ClassList) Add(class string) {
	c.Length++
	c.Value += " " + class
	c.Items = append(c.Items, class)
	c.object.Call("add", class)
}