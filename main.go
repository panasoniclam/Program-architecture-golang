package main

import (
	"fmt"
)

type FlagBit int

func main() {
	const (
		Send = 1 << iota
		Receive
		Contend
		Title
	)
	flag := Title
	fmt.Println(flag)
}
