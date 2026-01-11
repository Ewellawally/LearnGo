package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("你好" + "," + "世界")

	fmt.Println(quote.Go())
}
