package main

import (
	"fmt"
	"ltt"
)

func main() {
	fmt.Println(":c")

	new := ltt.Init()
	new.Write()
}
