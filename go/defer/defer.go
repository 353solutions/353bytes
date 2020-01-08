package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("defer")
	fmt.Println("working")
	fmt.Println("done")
}
