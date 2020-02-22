package ltt

import (
	"bytes"
	"time"
)

// specialChar holds all relevant information
// about any given character
type specialChar struct {
	// acutal character
	B string

	// occourences
	N int

	// percentage of all characters
	Perc float64
}

// key holds information about a key
type key struct {
	// of the length
	Len int

	// key's index
	Index int
}

// Report contains all informations about the analized keys
type Report struct {
	// time of the begin of the analysis
	initTime time.Time

	// Time difference as string
	PerformanceTime string

	// duration of the analysis
	AnalysisTime string

	// duration of the preparation
	PrepTime string

	// average length of a key
	Avr float64

	// number of key's scanned
	N int

	// shortest key by length
	Min key

	// longest key by length
	Max key

	// all duplicates among the keys
	duplicates bytes.Buffer

	// number of all duplicates
	nDuplicates int

	// number of each character occuring
	chars map[byte]int

	// number of all characters
	nChars int

	// most occouring character
	MostChar specialChar

	// least occouring character
	LeastChar specialChar
}
