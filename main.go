package main

import (
	"fmt"
	"go_learn/channels"
	"go_learn/structs"
)

func main() {
	fmt.Println("----channels----")
	channels.Example()
	fmt.Println("----structs----")
	a := getStringify()
	fmt.Println(a.GetString())
}

func getStringify() structs.Stringify {
	a := structs.Example()
	return &a
}
