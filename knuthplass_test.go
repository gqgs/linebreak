package linebreak

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestKnuthPlass runs table-driven tests with various inputs.
func Test_KnuthPlass(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		maxWidth int
		// expected output is provided when we know the exact line breaks.
		expected []string
	}{
		{
			name:     "Empty words",
			words:    []string{},
			maxWidth: 10,
			expected: []string{},
		},
		{
			name:     "Non-positive maxWidth",
			words:    []string{"Hello", "world"},
			maxWidth: 0,
			// Our implementation returns nil for non-positive maxWidth.
			expected: []string{},
		},
		{
			name:     "Single word",
			words:    []string{"Hello"},
			maxWidth: 10,
			expected: []string{"Hello"},
		},
		{
			name:  "Two words exact fit",
			words: []string{"Hello", "world"},
			// "Hello world" is 5 + 1 + 5 = 11 characters.
			maxWidth: 11,
			expected: []string{"Hello world"},
		},
		{
			name:  "Multiple words with wrapping",
			words: []string{"a", "b", "c", "d", "e"},
			// With maxWidth=3:
			// The optimal break is "a b", "c d", "e"
			maxWidth: 3,
			expected: []string{"a b", "c d", "e"},
		},
		{
			name:  "Each word exactly fits maxWidth",
			words: []string{"hello", "world"},
			// Each word is 5 letters and fits exactly.
			maxWidth: 5,
			expected: []string{"hello", "world"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := KnuthPlass(tc.words, tc.maxWidth)
			assert.Equal(t, tc.expected, result)
		})
	}
}

// TestKnuthPlassInvariants checks properties of the output even when the exact break points may vary.
func Test_KnuthPlassInvariants(t *testing.T) {
	words := []string{"This", "is", "a", "test", "of", "the", "emergency", "broadcast", "system"}
	maxWidth := 12

	lines := KnuthPlass(words, maxWidth)

	// Invariant 1: No line should exceed maxWidth.
	for i, line := range lines {
		if len(line) > maxWidth {
			t.Errorf("Line %d exceeds maxWidth: %q (length %d, maxWidth %d)", i, line, len(line), maxWidth)
		}
	}

	// Invariant 2: When splitting each line by spaces and reassembling, the original words appear in order.
	var rejoined []string
	for _, line := range lines {
		// Using strings.Fields to avoid issues with multiple spaces.
		rejoined = append(rejoined, strings.Fields(line)...)
	}

	assert.Equal(t, rejoined, words)
}
