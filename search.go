package ltt

import (
	"fmt"
	"time"
)

// Search for duplicates upon a slice of strings
func (r *Report) Search(objs []string) {
	start := time.Now()
	opTime := time.Now()
	length := len(objs)
	for i := 0; i < length; i++ {
		printProgressBar(i, length-1, opTime)
		opTime = time.Now()
		for j := i + 1; j < length; j++ {
			if objs[i] == objs[j] {
				if len(objs[i]) > 0 {
					r.AddDuplicate(objs[i], i+1)
				}
			}
		}
		r.Update(objs[i], i+1)
	}
	fmt.Println("")

	// Set analysis time
	r.AnalysisTime = time.Now().Sub(start).String()
}
