package slowmath

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

type SqrtCase struct {
	Value     float64
	Expected  float64
	ShouldErr bool `yaml:"error"`
}

func loadCases(t *testing.T, path string) []SqrtCase {
	require := require.New(t)
	file, err := os.Open(path)
	require.NoError(err)

	var cases []SqrtCase
	dec := yaml.NewDecoder(file)
	require.NoError(dec.Decode(&cases))
	return cases
}

func TestSqrt(t *testing.T) {
	testCases := loadCases(t, "sqrt_cases.yml")
	for _, tc := range testCases {
		name := fmt.Sprintf("%v", tc.Value)
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			out, err := Sqrt(tc.Value)

			if tc.ShouldErr {
				require.Error(err, "didn't get an error")
				return
			}

			require.NoError(err, "unexpected error")
			require.InDelta(tc.Expected, out, 0.001)
		})
	}
}
