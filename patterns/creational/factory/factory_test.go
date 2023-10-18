package factory

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactory(t *testing.T) {
	type testCase struct {
		Name           string
		Cast           string
		ExpectedResult string
		ExpectedErr    error
	}

	testCases := []*testCase{
		{
			Name:           "success whatsapp",
			Cast:           WHATSAPP_MESSENGER,
			ExpectedResult: fmt.Sprintf("%s %s", DEFAULT_MESSAGE, WHATSAPP_MESSENGER),
			ExpectedErr:    nil,
		},
		{
			Name:           "success telegram",
			Cast:           TELEGRAM_MESSENGER,
			ExpectedResult: fmt.Sprintf("%s %s", DEFAULT_MESSAGE, TELEGRAM_MESSENGER),
			ExpectedErr:    nil,
		},
		{
			Name:           "success viber",
			Cast:           VIBER_MESSENGER,
			ExpectedResult: fmt.Sprintf("%s %s", DEFAULT_MESSAGE, VIBER_MESSENGER),
			ExpectedErr:    nil,
		},
		{
			Name:           "success error handle",
			Cast:           "doesn't_exist_messenger",
			ExpectedResult: "",
			ExpectedErr:    ErrMessengerNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			messenger, err := MessengerFactory(tc.Cast)
			if err != nil {
				if !assert.Equal(t, tc.ExpectedErr, err) {
					t.Errorf("extected that error: %s, got that error: %s", tc.ExpectedErr, err)
				}
				return
			}

			got := messenger.Send()
			if !assert.Equal(t, tc.ExpectedResult, got) {
				t.Errorf("expected that: %s got that: %s", tc.ExpectedResult, got)
			}

		})
	}

}
