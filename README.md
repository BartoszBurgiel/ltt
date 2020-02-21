# Little Testing Tool 
Useful to test encryption

---

# LTT 
* Has dummy data 
    * 200-ish  (non-repeating) first names
    * 200-ish  (non-repeating) last names 
    * 1.600.000-ish (non-repeating) random words
* Analyse provided keys 
    * Search for duplicates among provided keys 
    * Keys need to be provided in a single file called `keys.txt`
    * each key seperated by `\n`
    * generate a pretty `report.txt` file with 
        * average length of all keys  
        * longest key 
        * shortest key
        * time to perform 
* It's very cool in general

### How to use
* `ltt <path-to-keys.txt>` - search for duplicates and return _nice_ log file
* `-gF` - generate firstNames.txt containing all 200ish first names seperated by \n
* `-gL` - generate lastNames.txt containing all 200ish last names seperated by \n
* `-gN` - generate names.txt containing all 40.000ish full names seperated by \n
* `-gW` - generate words.txt containing all 1.600.000ish words seperated by \n