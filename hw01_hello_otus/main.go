package main

import (
	"fmt"
	stUtil "golang.org/x/example/stringutil"
)

func main() {
	var name = "Hello, OTUS!"

	fmt.Println(stUtil.Reverse(name))
}
