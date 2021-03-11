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

	a, b := multi.vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := multi.vals()
	fmt.Println(c)
}
