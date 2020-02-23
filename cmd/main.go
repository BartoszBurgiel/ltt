package main

import (
	"fmt"
	"io/ioutil"
	"ltt"
	"ltt/dummy"
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

		// Divide to []string
		contentsStr := string(contents)
		str := strings.Split(contentsStr, "\n")

		if len(str) <= 1 {
			fmt.Println("To few keys in keys.txt")
			os.Exit(0)
		}

		fmt.Println("Analysing...")
		new := ltt.Init()
		new.Search(str)
		new.Write(str)
		os.Exit(0)
	}

	switch argument {
	case "-gF":
		// File with first names
		ltt.GenerateDummyNames("first", dummy.FirstNames)
		os.Exit(0)
	case "-gL":
		// File with last names
		ltt.GenerateDummyNames("last", dummy.Surnames)
		os.Exit(0)
	case "-gN":
		// File with full names
		ltt.GenerateFullNames(dummy.FirstNames, dummy.Surnames)
		os.Exit(0)
	case "-gW":
		// File with words
		ltt.GenerateWords(dummy.Words)
		os.Exit(0)
	case "-h":
		// Help
		fmt.Println(`* ltt <path-to-keys.txt> - search for duplicates and return _nice_ report.txt file
* -gF - generate firstNames.txt containing all 1000ish first names seperated by \n
* -gL - generate lastNames.txt containing all 1000ish last names seperated by \n
* -gN - generate names.txt containing all 1.000.000ish full names seperated by \n
* -gW - generate words.txt containing all 1.800.000ish words seperated by \n`)
	default:
		fmt.Println("Invalid argument")
		os.Exit(0)
	}
}
