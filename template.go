package ltt

// report template
var reportTemplate string = `Little Testing Tool's report

Total number of analysed keys: {{.N}}
Average length of the keys: {{.Avr}}

Longest key's length: {{.Max.Len}}
Longest key's position (line): {{.Max.Index}}

Shortest key's length: {{.Min.Len}}
Shortest key's position (line): {{.Min.Index}}

Most common character: '{{.MostChar.B}}' 
	- Number of occourences: {{.MostChar.N}}
	- (statistically) Number of occourences in every key: {{.MostChar.NWords}}
	- Percentage of all characters: {{.MostChar.Perc}}%
	
Least common character: '{{.LeastChar.B}}' 
	- Number of occourences: {{.LeastChar.N}} 
	- (statistically) Number of occourences in every key: {{.LeastChar.NWords}}
	- Percentage of all characters: {{.LeastChar.Perc}}%

Similarity: {{.Similarity}}%
	- Highest similarity of a position: {{.HPC.Similarity}}%
		- '{{.HPC.B}}' on position {{.HPC.Pos}} -> {{.HPC.Occ}} occourences
	- Lowest similarity of a position: {{.LPC.Similarity}}%
		- '{{.LPC.B}}' on position {{.LPC.Pos}} -> {{.LPC.Occ}} occourences
	- Standard abbreviation of the positions: {{.StandardAbbrvSim}}

Total performance time: {{.PerformanceTime}}
	- Analysis duration: {{.AnalysisTime}}
	- LTT's preperation: {{.PrepTime}}


`

// duplicates template for the report
var duplicatesHeader string = `--- DUPLICATES FOUND ---

Number of found duplicates: %d
Percantage of duplicates: %.4f%%

n    |pos    |
`

// duplicate template for the buffer
var duplicateTemplate string = "%-5d|%-7d|%s\n"
