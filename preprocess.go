package ltt

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

// Precalculate and prepare all relevant information
// about scanned keys (clean up the Write() method)
func (r *Report) prepare(keys []string) {
	start := time.Now()

	// GatherData for the report
	fmt.Println("\nAnalysing keys...")
	r.GatherData(keys)

	startCalculation := time.Now()
	fmt.Println("Calculating...")

	// Calculate average
	r.Avr /= float64(r.N)

	// Establish the most and least occouring
	// character and its occourings
	for b, occ := range r.chars {
		if occ > r.MostChar.N {
			r.MostChar.B = string(b)
			r.MostChar.N = occ
		}
		if occ < r.LeastChar.N && occ != 0 {
			r.LeastChar.B = string(b)
			r.LeastChar.N = occ
		}
		r.nChars += occ
	}

	// calculate percentages
	r.MostChar.Perc = float64(r.MostChar.N) / float64(r.nChars)
	r.LeastChar.Perc = float64(r.LeastChar.N) / float64(r.nChars)

	r.MostChar.NWords = float64(r.MostChar.N) / float64(r.N)
	r.LeastChar.NWords = float64(r.LeastChar.N) / float64(r.N)

	// calculate similarity index
	r.calculateSimilarities()

	// calculate the complexity of the keys
	r.calculateComplexity()

	// round the variables
	round(&r.Avr, &r.MostChar.Perc, &r.LeastChar.Perc, &r.LeastChar.NWords, &r.MostChar.NWords, &r.LC.Complexity, &r.HC.Complexity, &r.Complexity)

	// Format the integers
	r.format()

	r.CalculationTime = time.Now().Sub(startCalculation).String()

	// Add preperation duration
	r.PrepTime = time.Now().Sub(start).String()
}

// Format all integers that need to be formatted
// -> add commas
func (r *Report) format() {

	// formatting function
	form := func(n int) (out string) {

		// convert to string
		nn := strconv.Itoa(n)

		for i := 0; i < len(nn); i++ {
			out += string(nn[i])

			// at every 3rd add ','
			if (len(nn)-1-i)%3 == 0 && i != len(nn)-1 {
				out += ","
			}
		}
		return out
	}

	// Format the variables

	r.Nform = form(r.N)
	r.nDuplicatesForm = form(r.nDuplicates)

	r.HPC.OccForm = form(r.HPC.Occ)
	r.HPC.PosForm = form(r.HPC.Pos)

	r.LPC.OccForm = form(r.LPC.Occ)
	r.LPC.PosForm = form(r.LPC.Pos)

	r.LeastChar.Nform = form(r.LeastChar.N)
	r.MostChar.Nform = form(r.MostChar.N)

	r.Max.IndexForm = form(r.Max.Index)
	r.Min.IndexForm = form(r.Min.Index)

	r.LC.PositionForm = form(r.LC.Position)
	r.HC.PositionForm = form(r.HC.Position)
}

// custom rounding function
func round(n ...*float64) {
	for _, nn := range n {
		// round and write to address
		*nn = math.Round((*nn)*1000) / 1000
	}
}

// GatherData gathers data for the report
// that will be analysed later on
func (r *Report) GatherData(keys []string) {
	start := time.Now()
	opTime := time.Now()
	fmt.Println("Gathering additional data...")
	for i, key := range keys {

		// display progress bar
		printProgressBar(i, len(keys)-1, opTime)
		opTime = time.Now()
		// If key not empty
		if len(key) != 0 {
			r.Update(key, i)
		}
	}
	fmt.Println("")
	// set time of gathering compeltion
	r.GatheringDataTime = time.Now().Sub(start).String()
}
