package main

import (
	"fmt"
	"os"
)

// Bar bar
func Bar() string {
	return "bar"
}

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " and "
	}
	fmt.Println(s)
}
