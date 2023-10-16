package builder

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManufacturingDirector_Builder(t *testing.T) {
	type testCase struct {
		Name           string
		MD             ManufacturingDirector
		Cast           BuildProcess
		ExpectedResult Gun
	}

	testCases := []*testCase{
		{
			Name: "success pistol",
			MD:   ManufacturingDirector{},
			Cast: new(Pistol),
			ExpectedResult: Gun{
				Ammo:      10,
				Damage:    5,
				Structure: "Pistol",
			},
		},
		{
			Name: "success rifle",
			MD:   ManufacturingDirector{},
			Cast: new(Rifle),
			ExpectedResult: Gun{
				Ammo:      30,
				Damage:    10,
				Structure: "Rifle",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			manufacturingComplex := ManufacturingDirector{}

			manufacturingComplex.SetBuilder(tc.Cast)
			manufacturingComplex.Construct()
			got := tc.Cast.GetGun()

			if !assert.Equal(t, tc.ExpectedResult, got) {
				t.Errorf("expected resutl %v got that: %v", tc.ExpectedResult, got)
			}

			fmt.Printf("GOT: %v\n", got)
		})
	}
}
