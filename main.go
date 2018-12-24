package main

import (
	. "./goquery"
	"fmt"
)


func main() {
	main, _ := Q("#main")
	fmt.Printf("class list: %s\n", main.ClassList.Value)
	main.ClassList.Add("quux")
	fmt.Printf("added class quux, new class list: %s\n", main.ClassList.Value)
}