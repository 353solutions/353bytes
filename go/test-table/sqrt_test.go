package slowmath

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSqrt(t *testing.T) {
	var testCases = []struct {
		val       float64
		expected  float64
		shouldErr bool
	}{
		{2.0, 1.4142, false},
		{0.0, 0.0, false},
		{-1.0, 0.0, true},
	}
	for _, tc := range testCases {
		name := fmt.Sprintf("%v", tc.val)
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			out, err := Sqrt(tc.val)

			if tc.shouldErr {
				require.Error(err, "didn't get an error")
				return
			}

			require.NoError(err, "unexpected error")
			require.InDelta(tc.expected, out, 0.001)
		})
	}
}
