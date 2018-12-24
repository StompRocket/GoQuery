package goquery

import (
	"syscall/js"
)

type Event int

const (
	Click Event = iota
)

func (d *DomObject) On(event Event, callback func([]js.Value)) {

	switch event {
	case Click:
		d.object.Set("onclick", js.NewCallback(callback))
	}

}
