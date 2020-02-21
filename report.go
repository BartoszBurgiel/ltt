package ltt

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

// Report contains all informations about the analized keys
type Report struct {
	// time of the begin of the analysis
	initTime time.Time

	// average length of a key
	avr float64

	// number of key's scanned
	n int

	// shortest key by length
	min struct {

		// actual length
		len int

		// key's index
		index int
	}

	// longest key by length
	max struct {
		// actual length
		len int

		// key's index
		index int
	}

	// all duplicates among the keys
	duplicates bytes.Buffer

	// number of all duplicates
	nDuplicates int
}

// Update Report struct
// check if the provided key
// is the shortest or longest
// and update the avr length
func (r *Report) Update(key string, index int) {
	length := len(key)
	r.n++
	r.avr += float64(length) / float64(r.n)
	if length > r.max.len {
		r.max.len = length
		r.max.index = index
		return
	}
	if length < r.min.len {
		r.min.len = length
		r.min.index = index
		return
	}
}

// Write the report to report.txt file
func (r Report) Write() {
	// create report.txt
	file, err := os.Create("report.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()

	// Calculate the time difference
	diff := time.Now().Sub(r.initTime)

	// Compose report file contents with the template
	content := fmt.Sprintf(template, r.n, r.avr, r.max.len, r.max.index, r.min.len, r.min.index, diff.String())

	// Write content to report.txt
	file.WriteString(content)
	file.Sync()

	// Check for duplicates
	if r.duplicates.Len() == 0 {
		file.WriteString("No duplicates found :)")
		file.Sync()
		os.Exit(0)
	}

	// prepare the duplicate header
	percentage := 100.0 * float64(r.nDuplicates) / float64(r.n)
	content = fmt.Sprintf(duplicatesHeader, r.nDuplicates, percentage)

	// Write duplicate header
	file.WriteString(content)

	file.Write(r.duplicates.Bytes())
	file.Sync()
}

// Init ialize the report
func Init() Report {
	return Report{
		initTime: time.Now(),
		min: struct {
			len   int
			index int
		}{
			len:   999999999999999,
			index: 0,
		},
	}
}
