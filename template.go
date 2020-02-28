package ltt

// report template
var reportTemplate string = `Little Testing Tool's report 

File: {{.FileName}}
Date: {{.AnalysisDate}}

Total number of analysed keys: {{.Nform}}
Average length of the keys: {{.Avr}}

Longest key's length: {{.Max.Len}}
Longest key's position (line): {{.Max.IndexForm}}

Shortest key's length: {{.Min.Len}}
Shortest key's position (line): {{.Min.IndexForm}}

Most common character: '{{.MostChar.B}}' 
	- Number of occourences: {{.MostChar.Nform}}
	- (statistically) Number of occourences in every key: {{.MostChar.NWords}}
	- Percentage of all characters: {{.MostChar.Perc}}%
	
Least common character: '{{.LeastChar.B}}' 
	- Number of occourences: {{.LeastChar.Nform}} 
	- (statistically) Number of occourences in every key: {{.LeastChar.NWords}}
	- Percentage of all characters: {{.LeastChar.Perc}}%

Similarity among the keys: {{.Similarity}}%
	- Highest similarity of a position: {{.HPC.Similarity}}%
		- '{{.HPC.B}}' on position {{.HPC.PosForm}} with {{.HPC.OccForm}} occourences
	- Lowest similarity of a position: {{.LPC.Similarity}}%
		- '{{.LPC.B}}' on position {{.LPC.PosForm}} with {{.LPC.OccForm}} occourences
	- Standard deviation of the similarities: {{.StandardDevSim}} [%]

Average complexity of the keys: {{.Complexity}}%
	- Highest complexity of a key: {{.HC.Complexity}}%
		- On position {{.HC.PositionForm}}
	- Lowest complexity of a key: {{.LC.Complexity}}%
		- On position {{.LC.PositionForm}}

Total performance time: {{.PerformanceTime}}
	- Duplicates search: {{.AnalysisTime}}
	- LTT's preperation: {{.PrepTime}}
		- Data gathering: {{.GatheringDataTime}}
		- Calculation Time: {{.CalculationTime}}


`

// duplicates template for the report
var duplicatesHeader string = `--- DUPLICATES FOUND ---

Number of found duplicates: %d
Percantage of duplicates: %.4f%%

n    |pos    |
`

// duplicate template for the buffer
var duplicateTemplate string = "%-5d|%-7d|%s\n"
