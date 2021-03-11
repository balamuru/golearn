package main

import (
	"fmt"
	"github.com/balamuru/golearn/func/multi"

)

func main() {

	a, b := multi.Vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := multi.Vals()
	fmt.Println(c)
}
