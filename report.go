package ltt

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

// Update Report struct
// check if the provided key
// is the shortest or longest
// and update the avr length
func (r *Report) Update(key string, index int) {

	// Add values to chars
	for i, b := range key {
		r.chars[byte(b)]++

		// If current position not existant -> add
		if len(r.positions) <= i {
			r.positions = append(r.positions, pos{})
		}
		r.positions[i][byte(b)]++
	}

	length := len(key)
	r.N++
	r.Avr += float64(length)
	if length > r.Max.Len {
		r.Max.Len = length
		r.Max.Index = index
		return
	}
	if length < r.Min.Len && length > 0 {
		r.Min.Len = length
		r.Min.Index = index
		return
	}
}

// AddDuplicate to the duplicate buffer
func (r *Report) AddDuplicate(key string, index int) {
	r.nDuplicates++

	// prepare string
	prep := fmt.Sprintf(duplicateTemplate, r.nDuplicates, index, key)
	r.duplicates.WriteString(prep)
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

	// Precalculate the numbers
	r.prepare()

	// Calculate the time difference
	r.PerformanceTime = time.Now().Sub(r.initTime).String()

	tmpl, err := template.New("test").Parse(reportTemplate)

	// Compose report file contents with the template
	err = tmpl.Execute(file, r)
	file.Sync()

	// Check for duplicates
	if r.duplicates.Len() == 0 {
		file.WriteString("No duplicates found :)")
		file.Sync()
		os.Exit(0)
	}

	// prepare the duplicate header
	percentage := 100.0 * float64(r.nDuplicates) / float64(r.N)
	content := fmt.Sprintf(duplicatesHeader, r.nDuplicates, percentage)

	// Write duplicate header
	file.WriteString(content)

	file.Write(r.duplicates.Bytes())
	file.Sync()
}

// Init ialize the report
func Init() Report {
	return Report{
		initTime: time.Now(),
		Min: key{
			Len:   999999999999999,
			Index: 0,
		},
		chars: make(map[byte]int),
		LeastChar: specialChar{
			N: 999999999999999,
		},
		LPC: similarityChar{
			Occ: 99999999999999,
		},
	}
}
