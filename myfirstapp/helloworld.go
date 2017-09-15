package main;

import (
	"fmt";
	"mastering_go/myfirstapp/utility"
)

func main() {
	fmt.Println("Hello, World!");
	utility.SayHello();
	// Won't work because funcitons with begining with lower case are private.
	// utility.sayMyName();
}
