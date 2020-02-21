package main

import (
	"fmt"
	"io/ioutil"
	"ltt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("No arguments provided")
		os.Exit(0)
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments provided")
		os.Exit(0)
	}

	argument := os.Args[1]

	if argument[0] != '-' {
		fmt.Println("Reading file's contents...")
		// Read file's content
		contents, err := ioutil.ReadFile(argument)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		contentsStr := string(contents)
		str := strings.Split(contentsStr, "\n")

		fmt.Println("Analysing...")
		new := ltt.Init()
		new.Search(str)
		new.Write()
		os.Exit(0)
	}
}
