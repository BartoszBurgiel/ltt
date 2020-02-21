package ltt

import (
	"bytes"
	"fmt"
	"os"
)

// GenerateDummyNames and write the names
// provided in the parameter to the file
func GenerateDummyNames(prefix string, names []string) {
	// create [prefix]Names.txt
	file, err := os.Create(prefix + "Names.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()

	temp := bytes.Buffer{}

	for _, name := range names {
		temp.WriteString(name + "\n")

		// control the size
		if temp.Len() > 300000 {
			file.Write(temp.Bytes())
			file.Sync()
			temp.Reset()
		}
	}

	// Write leftovers
	file.Write(temp.Bytes())
	file.Sync()
	temp.Reset()

	fmt.Println(len(names), prefix, "names have been written")
}

// GenerateFullNames out of given first and last names and
// write them to the file
func GenerateFullNames(firstNames, lastNames []string) {
	// create names.txt
	file, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()

	temp := bytes.Buffer{}

	for _, firstName := range firstNames {
		for _, lastName := range lastNames {

			temp.WriteString(firstName + " " + lastName + "\n")

			// control the size
			if temp.Len() > 300000 {
				file.Write(temp.Bytes())
				file.Sync()
				temp.Reset()
			}
		}
	}

	// Write leftovers
	file.Write(temp.Bytes())
	file.Sync()
	temp.Reset()
	fmt.Println(len(firstNames)*len(lastNames), "names have been generated and written")
}

// GenerateWords writes words to the file
func GenerateWords(words string) {
	// create words.txt
	file, err := os.Create("words.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()

	file.WriteString(words)
	file.Sync()

	// count words
	n := 0
	for _, b := range words {
		if b == '\n' {
			n++
		}
	}

	fmt.Println(n, "words have been written")
}
