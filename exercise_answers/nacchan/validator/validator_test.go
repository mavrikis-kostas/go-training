package validator

import (
	"testing"
)

func TestIsInRange(t *testing.T) {
	tt := []struct {
		name          string
		num, min, max int
		expected      bool
	}{
		{
			name:     "within range",
			num:      5,
			min:      1,
			max:      10,
			expected: true,
		},
		{
			name:     "below range",
			num:      0,
			min:      1,
			max:      10,
			expected: false,
		},
		{
			name:     "above range",
			num:      11,
			min:      1,
			max:      10,
			expected: false,
		},
		{
			name:     "equal to min",
			num:      1,
			min:      1,
			max:      10,
			expected: true,
		},
		{
			name:     "equal to max",
			num:      10,
			min:      1,
			max:      10,
			expected: true,
		},
		{
			name:     "zero",
			num:      0,
			min:      0,
			max:      0,
			expected: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := IsInRange(tc.num, tc.min, tc.max)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}
