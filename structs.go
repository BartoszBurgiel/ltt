package ltt

import (
	"bytes"
	"time"
)

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

	// all of the positions
	positions []pos

	// Similarity index
	Similarity float64

	// Highest similarity char with it's position
	HPC similarityChar

	// Lowest similarity char with it's position
	LPC similarityChar

	// Standart Abbriviation upon the similarities
	StandardAbbrvSim float64
}

// specialChar holds all relevant information
// about any given character
type specialChar struct {
	// acutal character
	B string

	// occourences
	N int

	// statistical occourences per word
	NWords float64

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

// pos will be an array where each index
// represents a position of a key (index)
// it stores all occourences of the bytes
type pos map[byte]int

// stores all information about a char
// at certain position with certain similarity
type similarityChar struct {
	// the char
	B string

	// position
	Pos int

	// actual similarity
	Similarity float64

	// total occourences of the char
	Occ int
}
