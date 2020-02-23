package ltt

import (
	"math"
	"strconv"
	"time"
)

// Precalculate and prepare all relevant information
// about scanned keys (clean up the Write() method)
func (r *Report) prepare() {
	start := time.Now()

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

	// round the variables
	round(&r.Avr, &r.MostChar.Perc, &r.LeastChar.Perc, &r.LeastChar.NWords, &r.MostChar.NWords)

	// calculate similarity index
	r.calculateSimilarities()

	// Format the integers
	r.format()

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
}

func round(n ...*float64) {
	for _, nn := range n {
		*nn = math.Round((*nn)*1000) / 1000
	}
}
