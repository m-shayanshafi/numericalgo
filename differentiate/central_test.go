package differentiate_test

import (
	"math"
	"testing"

	"github.com/DzananGanic/numericalgo/differentiate"
	"github.com/stretchr/testify/assert"
)

func TestCentral(t *testing.T) {

	cases := map[string]struct {
		f             func(x float64) float64
		val           float64
		h             float64
		expectedValue float64
		expectedError error
	}{
		"central difference with 0.1 step size": {
			f: func(x float64) float64 {
				return math.Cos(math.Pow(x, 2) - 2)
			},
			val:           1,
			h:             0.1,
			expectedValue: 1.6609,
			expectedError: nil,
		},
		"central difference with 0.01 step size": {
			f: func(x float64) float64 {
				return math.Cos(math.Pow(x, 2) - 2)
			},
			val:           1,
			h:             0.01,
			expectedValue: 1.6827,
			expectedError: nil,
		},
		"central difference with 0.001 step size": {
			f: func(x float64) float64 {
				return math.Cos(math.Pow(x, 2) - 2)
			},
			val:           1,
			h:             0.001,
			expectedValue: 1.6829,
			expectedError: nil,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := differentiate.Central(c.f, c.val, c.h)
			assert.InEpsilon(t, result, c.expectedValue, 1e-4)
			assert.Equal(t, err, c.expectedError)
		})
	}
}
