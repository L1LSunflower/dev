package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToMap(t *testing.T) {
	type testCase struct {
		Name           string
		Cast           Adapter
		ExpectedResult map[string]any
	}

	testCases := []testCase{
		{
			Name: "success request",
			Cast: &Request{
				username:   "test_username",
				password:   "test_password",
				permission: 1,
			},
			ExpectedResult: map[string]any{
				"username":   "test_username",
				"password":   "test_password",
				"permission": 1,
			},
		},
		{
			Name: "success response",
			Cast: &Response{
				id:         1,
				token:      "test_token",
				permission: 1,
			},
			ExpectedResult: map[string]any{
				"id":         1,
				"token":      "test_token",
				"permission": 1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			got := tc.Cast.ToMap()
			if !assert.Equal(t, tc.ExpectedResult, got) {
				t.Errorf("expected that: %v got that: %v", tc.ExpectedResult, got)
			}
		})
	}
}
