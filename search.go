package ltt

// Search for duplicates upon a slice of strings
func Search(objs []string) {
	for i := 0; i < len(objs); i++ {
		for j := i + 1; j < len(objs); j++ {
			if objs[i] == objs[j] {
			}
		}
	}
}
