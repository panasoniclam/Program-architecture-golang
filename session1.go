package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "word!"
	if len(os.Args) > 0 {
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("hello", who)
}
