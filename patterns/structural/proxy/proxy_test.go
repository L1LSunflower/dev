package proxy

import "testing"

func TestProxy_GetAccount(t *testing.T) {
	type testCase struct {
		Name string
		Cast struct {
			id       string
			username string
			password string
		}
		ExpectedResult Account
	}

	testCases := []*testCase{
		{
			Name: "",
			Cast: struct {
				id       string
				username string
				password string
			}{
				id:       "",
				username: "",
				password: "",
			},
			ExpectedResult: Account{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

		})
	}
}
