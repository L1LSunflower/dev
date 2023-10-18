package builder

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManufacturingDirector_Builder(t *testing.T) {
	type testCase struct {
		Name           string
		MD             ManufacturingDirector
		Cast           BuildProcess
		ExpectedResult struct {
			Method string
			Url    string
			Body   []byte
		}
	}

	testCases := []*testCase{
		{
			Name: "success post request",
			MD:   ManufacturingDirector{},
			Cast: new(HttpPost),
			ExpectedResult: struct {
				Method string
				Url    string
				Body   []byte
			}{Method: http.MethodPost, Url: "http://test_method_post.com", Body: []byte("{\"data\": \"hello_post\"}")},
		},
		{
			Name: "success get request",
			MD:   ManufacturingDirector{},
			Cast: new(HttpGet),
			ExpectedResult: struct {
				Method string
				Url    string
				Body   []byte
			}{Method: http.MethodGet, Url: "http://test_method_get.com", Body: []byte("{\"data\": \"hello_get\"}")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			expectedRequest, _ := http.NewRequest(tc.ExpectedResult.Method, tc.ExpectedResult.Url, bytes.NewReader(tc.ExpectedResult.Body))

			manufacturingComplex := ManufacturingDirector{}

			manufacturingComplex.SetBuilder(tc.Cast)
			manufacturingComplex.Construct()
			got := tc.Cast.GetRequest()

			if !assert.Equal(t, expectedRequest.Method, got.Method) {
				t.Errorf("expected result %v got that: %v", expectedRequest.Method, got.Method)
			}

			if !assert.Equal(t, expectedRequest.URL.String(), got.URL.String()) {
				t.Errorf("expected result %v got that: %v", expectedRequest.URL.String(), got.URL.String())
			}

			if !assert.Equal(t, expectedRequest.Body, got.Body) {
				t.Errorf("expected result %v got that: %v", expectedRequest.Body, got.Body)
			}
		})
	}
}
