package main

import (
	"fmt"
)

func main() {
	x := 7
	changeX(&x)
	fmt.Println(x)
}

func changeX(arg *int) {
	*arg = 8
}
