package main

// import "fmt"

import (
	"fmt"
	"github.com/balamuru/golearn/func/multi"

)

//TODO delete
// func vals() (int, int) {
// 	return 3, 7
// }

func main() {

	a, b := multi.Vals
	fmt.Println(a)
	fmt.Println(b)

	_, c := multi.Vals()
	fmt.Println(c)
}
