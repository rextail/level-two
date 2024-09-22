package tests_test

import (
	"dev05/config"
	"dev05/internal/grep"
	"testing"
)

func TestGrepABC(t *testing.T) {
	tests := []struct {
		name     string
		indexes  []int
		count    int
		expected [][2]int
	}{
		{
			name:     "No Overlap",
			indexes:  []int{1, 4, 8},
			count:    1,
			expected: [][2]int{{1, 2}, {4, 5}, {8, 9}},
		},
		{
			name:     "With Overlap",
			indexes:  []int{1, 2, 4, 5},
			count:    1,
			expected: [][2]int{{1, 3}, {4, 6}},
		},
		{
			name:     "Single Index",
			indexes:  []int{2},
			count:    2,
			expected: [][2]int{{2, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.Config{After: tt.count}
			grepper := grep.New(cfg)

			result := grepper.GrepABC(tt.indexes)
			if len(result) != len(tt.expected) {
				t.Errorf("expected %v ranges, got %v", len(tt.expected), len(result))
			}

			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("expected range %v, got %v", tt.expected[i], result[i])
				}
			}
		})
	}
}
