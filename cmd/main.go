package main

import (
	"fmt"
	"ltt"
)

func main() {
	fmt.Println("Analysing...")

	new := ltt.Init()

	new.Update("hehe", 1)
	new.Update("hee", 2)
	new.Update("heheqwe", 3)
	new.Update("heh", 4)
	new.Update("hesddsadsahe", 5)
	new.Write()
}
