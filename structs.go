package ltt

import (
	"bytes"
	"time"
)

// Report contains all informations about the analized keys
type Report struct {

	// name of the analysing file
	FileName string

	// date of the analysis
	AnalysisDate string

	// time of the begin of the analysis
	initTime time.Time

	// Time difference as string
	PerformanceTime string

	// duration of the analysis
	AnalysisTime string

	// duration of the preparation
	PrepTime string

	// duration of the data gathering
	GatheringDataTime string

	// duration of the calculation
	CalculationTime string

	// average length of a key
	Avr float64

	// number of key's scanned
	N int

	// Formated N
	Nform string

	// shortest key by length
	Min key

	// longest key by length
	Max key

	// all duplicates among the keys
	duplicates bytes.Buffer

	// number of all duplicates
	nDuplicates int

	// Formatted nDuplicates
	nDuplicatesForm string

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

	// Standart Deviation upon the similarities
	StandardDevSim float64

	// Key with the highest complexity
	HC complexity

	// Key with the lowest complexity
	LC complexity

	// Average complexity of all keys
	Complexity float64

	// Complexity map
	com com
}

// specialChar holds all relevant information
// about any given character
type specialChar struct {
	// acutal character
	B string

	// occourences
	N int

	// Formatted N
	Nform string

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

	// key's index formatted
	IndexForm string
}

// pos will be an array where each index
// represents a position of a key (index)
// it stores all occourences of the bytes
type pos map[byte]int

// com represents the complexity of each key
// map's keys represents the keys' indexes
// map's values hold the data about every key
type com map[int]*complexityKey

// stores all information about a char
// at certain position with certain similarity
type similarityChar struct {
	// the char
	B string

	// position
	Pos int

	// Formatted position
	PosForm string

	// actual similarity
	Similarity float64

	// total occourences of the char
	Occ int

	// Formatted Occ
	OccForm string
}

// complexityKey stores all data needed to calculate
// the complexity of a key
type complexityKey struct {

	// length of the key
	len int

	// number of occurences of every character
	chars map[rune]int
}

// stores information about the general complexity of a
// certain key
type complexity struct {
	// Complexity value
	Complexity float64

	// Position
	Position int

	// Formatted position
	PositionForm string
}
