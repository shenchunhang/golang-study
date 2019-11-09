package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "3"
	}
	fmt.Println(s)
	fmt.Println(os.Args)
}
