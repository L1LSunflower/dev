package singleton

import (
	"testing"
)

func TestSingleton_AddChar(t *testing.T) {
	type testCase struct {
		Name           string
		Cast           string
		ExpectedResult string
	}

	testCases := []*testCase{
		{
			Name: "success get singleton",
		},
		{
			Name:           "success insert values singleton",
			Cast:           "abc",
			ExpectedResult: "abc",
		},
		{
			Name:           "success singleton with values",
			ExpectedResult: "abc",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Первый вызов для инициализации
			got := ReadChars()

			for _, b := range tc.Cast {
				singletonOnce.AddChar(byte(b))
			}

			// Перезаписываем значение
			got = ReadChars()
			if string(got) != tc.ExpectedResult {
				t.Fatalf("expected that: %v got that %v", tc.ExpectedResult, got)
			}

		})
	}

}
