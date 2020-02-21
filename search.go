package ltt

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

// search for duplicates upon a slice of strings
func search(objs []string) {
	for i := 0; i < len(objs); i++ {
		for j := i + 1; j < len(objs); j++ {
			if objs[i] == objs[j] {
			}
		}
	}
}

func main() {

	// Read file
	file, _ := ioutil.ReadFile(os.Args[1])
	strFile := string(file)

	avrLen := 0.0
	serpStrFile := strings.Split(strFile, "\n")
	for i := 0; i < len(serpStrFile); i++ {
		for j := i + 1; j < len(serpStrFile); j++ {
			if serpStrFile[i] == serpStrFile[j] {
				fmt.Println("duplicate!!!")
			}
		}
		avrLen += float64(len(serpStrFile[i]))
	}
	fmt.Println("no duplicates")
	fmt.Println("checked keys", len(serpStrFile))
	fmt.Println("avr len", avrLen/float64(len(serpStrFile)))

	fmt.Println(math.Pow(635.0, 85.0))
}
