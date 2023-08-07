package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	str := "Hello, OTUS!"
	strReversed := reverse.String(str)

	fmt.Println(strReversed)
}
