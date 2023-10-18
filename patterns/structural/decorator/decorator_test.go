package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Pay(t *testing.T) {
	type testCase struct {
		Name string
		Cast struct {
			amount float64
			sum    float64
		}
		ExpectedResult string
	}

	testCases := []testCase{
		{
			Name: "success",
			Cast: struct {
				amount float64
				sum    float64
			}{amount: 100, sum: 10},
			ExpectedResult: PaymentSuccess,
		},
		{
			Name: "fail not enough amount",
			Cast: struct {
				amount float64
				sum    float64
			}{amount: 10, sum: 10},
			ExpectedResult: PaymentFail,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			newPS := NewPSWithCommission(tc.Cast.amount)
			got := newPS.Pay(tc.Cast.sum)
			if !assert.Equal(t, tc.ExpectedResult, got) {
				t.Errorf("expected that: %s got that: %s", tc.ExpectedResult, got)
			}
		})
	}
}
