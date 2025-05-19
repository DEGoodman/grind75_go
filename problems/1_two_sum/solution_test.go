package problem1

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.nums, tt.target)
			if len(result) != len(tt.expected) {
				t.Errorf("got length %d, want length %d", len(result), len(tt.expected))
				return
			}
			// Add more specific assertions based on the problem requirements
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("twoSum(%v, %d) = %v, want %v", tt.nums, tt.target, result, tt.expected)
					return
				}
			}
		})
	}
}
