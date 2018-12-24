package main

import (
	. "./goquery"
	"fmt"
	"syscall/js"
)


func main() {
	main, _ := Q("#main")
	fmt.Printf("class list: %s\n", main.ClassList.Value)
	main.ClassList.Add("quux")
	fmt.Printf("added class quux, new class list: %s\n", main.ClassList.Value)

	main.On(Click, func(_ []js.Value) {
		fmt.Println("Clicked on #main")
	})

	done := make(chan struct{}, 0)
	<-done
}