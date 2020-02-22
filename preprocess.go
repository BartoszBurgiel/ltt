package ltt

import (
	"math"
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

	// round the variables
	round(&r.Avr, &r.MostChar.Perc, &r.LeastChar.Perc)

	// Add preperation duration
	r.PrepTime = time.Now().Sub(start).String()
}

func round(n ...*float64) {
	for _, nn := range n {
		*nn = math.Round((*nn)*1000) / 1000
	}
}
