package ltt

// report template
var template string = `Little Testing Tool's report

Total number of analysed keys: %d
Average length of the keys: %.5f

Longest key's length: %d
Longest key's position (line): %d

Shortest key's length: %d
Shortest key's position (line): %d

Performance time: %s

`

// duplicates template for the report
var duplicatesHeader string = `--- DUPLICATES FOUND ---

Number of found duplicates: %d
Percantage of duplicates: %.4f%%

n    |pos    |
`

// duplicate template for the buffer
var duplicateTemplate string = "%-5d|%-7d|%s\n"
