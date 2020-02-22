# Little Testing Tool 
Useful to test encryption

---

# LTT 
* Has dummy data 
    * credits https://github.com/smashew/NameDatabases
    * 1000-ish  (non-repeating) first names
    * 1000-ish  (non-repeating) last names 
    * 1.800.000-ish (non-repeating) random words
* Analyse provided keys 
    * Search for duplicates among provided keys 
    * Keys need to be provided in a single file called `keys.txt`
    * each key seperated by `\n`
    * generate a pretty `report.txt` file with 
        * average length of all keys  
        * longest key 
        * shortest key
        * most common character 
        * least common character
        * similarity among the keys
        * time to perform 
* It's very cool in general

# Similarity calculation 
Similarity calculator takes the average of the most commonly occouring letters on each position of the key
<a href="https://www.codecogs.com/eqnedit.php?latex=\sum_{i=0}^{n}&space;\frac{highestOccourance_i}{n}" target="_blank"><img src="https://latex.codecogs.com/gif.latex?\sum_{i=0}^{n}&space;\frac{highestOccourance_i}{n}" title="\sum_{i=0}^{n} \frac{highestOccourance_i}{n}" /></a> 

### How to use
* `ltt <path-to-keys.txt>` - search for duplicates and return _nice_ log file
* `-gF` - generate firstNames.txt containing all 1000ish first names seperated by \n
* `-gL` - generate lastNames.txt containing all 1000ish last names seperated by \n
* `-gN` - generate names.txt containing all 1.000.000ish full names seperated by \n
* `-gW` - generate words.txt containing all 1.800.000ish words seperated by \n