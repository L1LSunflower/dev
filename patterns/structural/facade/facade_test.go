package facade

import "testing"

func TestFacade(t *testing.T) {
	type testCase struct {
		Name           string
		Cast           string
		ExpectedResult string
	}

	testCases := []*testCase{
		{
			Name:           "",
			Cast:           "",
			ExpectedResult: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

		})
	}
}
