package ltt

func (r *Report) calculateSimilarities() {
	// stores the max percentage of each position
	percentages := []float64{}

	// Iterate over each positions
	for i, pos := range r.positions {

		max := 0
		// find the highest occourance
		for b, p := range pos {
			if p > max {
				max = p
			}

			// Set most and least occouring character on a position
			if p > r.HPC.Occ {
				r.HPC.Occ = p
				r.HPC.B = string(b)
				r.HPC.Pos = i
			}
			if p < r.LPC.Occ {
				r.LPC.Occ = p
				r.LPC.B = string(b)
				r.LPC.Pos = i
			}
		}
		// add to percentages
		percentages = append(percentages, float64(max)/float64(r.N))
	}

	// Calculate the average of percentages
	// and find highest and lowest percentage
	hi, lo := 0.0, 1.0
	avr := 0.0
	for _, p := range percentages {
		avr += p

		if p > hi {
			hi = p
		}
		if p < lo && p != 0 {
			lo = p
		}
	}

	avr /= float64(len(percentages))

	// Calculate the standart abbriviation
	sA := 0.0
	for _, p := range percentages {
		sA += abs(avr - p)
	}
	sA /= float64(len(percentages))

	r.Similarity = avr * 100
	r.LPC.Similarity = lo * 100
	r.HPC.Similarity = hi * 100
	r.StandardAbbrvSim = sA * 100
	round(&r.Similarity, &r.LPC.Similarity, &r.HPC.Similarity, &r.StandardAbbrvSim)
}

func abs(n float64) float64 {
	if n > 0 {
		return n
	}
	return -n
}
