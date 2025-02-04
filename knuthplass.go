package linebreak

import (
	"math"
	"strings"
)

// KnuthPlass wraps the text (provided as a slice of words) into lines
// not exceeding maxWidth. It uses a dynamic programming approach where the
// penalty for a line (except the last line) is defined as the cube of the extra spaces.
func KnuthPlass(words []string, maxWidth int) []string {
	// Validate maxWidth. Return an empty slice if the provided maxWidth is non-positive.
	if maxWidth <= 0 {
		return []string{}
	}

	n := len(words)
	// If there are no words, return an empty slice.
	if n == 0 {
		return []string{}
	}

	// cumLengths[i] stores the total length of words[0:i]
	cumLengths := make([]int, n+1)
	for i, word := range words {
		cumLengths[i+1] = cumLengths[i] + len(word)
	}

	// minCost[j] will store the minimum cost to break words[0:j]
	minCost := make([]float64, n+1)
	// lineBreaks[j] stores the index i that gives the best break for words[i:j]
	lineBreaks := make([]int, n+1)

	// Initialize minCost array with +Inf as the worst cost, except for the 0-th position.
	for i := 0; i <= n; i++ {
		minCost[i] = math.Inf(1)
	}
	minCost[0] = 0

	// Compute the dynamic programming states: for each possible break from word i to word j.
	for i := 0; i < n; i++ {
		// If current cost is infinite, skip processing.
		if math.IsInf(minCost[i], 1) {
			continue
		}
		for j := i + 1; j <= n; j++ {
			// Total length of words[i:j]
			wordsLen := cumLengths[j] - cumLengths[i]
			// Spaces between words (one space between each word)
			spaces := j - i - 1
			totalLen := wordsLen + spaces

			// If the total length exceeds maxWidth, no further words can be added.
			if totalLen > maxWidth {
				break
			}

			// Calculate the slack (extra spaces)
			slack := maxWidth - totalLen
			var cost float64
			// Do not penalize the last line.
			if j == n {
				cost = 0
			} else {
				// Use cubic penalty to discourage ragged lines.
				cost = float64(slack * slack * slack)
			}

			// Update the DP state if this break yields a lower cost.
			if minCost[i]+cost < minCost[j] {
				minCost[j] = minCost[i] + cost
				lineBreaks[j] = i
			}
		}
	}

	// Reconstruct the line break indices by walking backwards from the end.
	var breakIndices []int
	for j := n; j > 0; {
		breakIndices = append(breakIndices, j)
		j = lineBreaks[j]
	}
	breakIndices = append(breakIndices, 0) // include the starting index

	// Reverse breakIndices to obtain the correct order.
	for i, j := 0, len(breakIndices)-1; i < j; i, j = i+1, j-1 {
		breakIndices[i], breakIndices[j] = breakIndices[j], breakIndices[i]
	}

	// Build the final lines from the break indices.
	lines := make([]string, 0, len(breakIndices)-1)
	for i := 0; i < len(breakIndices)-1; i++ {
		start, end := breakIndices[i], breakIndices[i+1]
		line := strings.Join(words[start:end], " ")
		lines = append(lines, line)
	}

	return lines
}
