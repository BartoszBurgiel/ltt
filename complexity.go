package ltt

// Calculate how complex are the keys
func (r *Report) calculateComplexity() {

	// Complexity value
	comp := 0.0

	// iterate over complexity map
	for i, n := range r.com {

		// find the highest occourance of a char
		hi := 0
		for _, occ := range n.chars {
			if occ > hi {
				hi = occ
			}
		}

		// complexity value in percent
		comp = (float64(hi) / float64(n.len)) * 100

		// Add to overall complexity
		r.Complexity += comp

		// Establish the highest complexity
		if comp > r.HC.Complexity {
			r.HC.Complexity = comp
			r.HC.Position = i
		}

		// Establish the lowest complexity
		if comp < r.LC.Complexity {
			r.LC.Complexity = comp
			r.LC.Position = i
		}

	}

	// Divide by N to get the average
	r.Complexity /= float64(r.N)
}
