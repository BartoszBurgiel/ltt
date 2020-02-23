package ltt

import (
	"fmt"
	"time"
)

// Search for duplicates upon a slice of strings
func (r *Report) Search(objs []string) {
	// Start time of the analysis algorithm
	start := time.Now()

	// Operation time for the progress bar
	opTime := time.Now()

	// Precalculated length of the key slice
	length := len(objs)

	for i := 0; i < length; i++ {
		// Print progress bar and reset the time
		printProgressBar(i, length-1, opTime)
		opTime = time.Now()

		// linear search -> O(n^2)
		for j := i + 1; j < length; j++ {
			if objs[i] == objs[j] {
				if len(objs[i]) > 0 {
					r.AddDuplicate(objs[i], i+1)
				}
			}
		}
	}
	fmt.Println("")

	// Set analysis time
	r.AnalysisTime = time.Now().Sub(start).String()
}
