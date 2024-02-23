package cronparser_test

import (
	"testing"

	"github.com/KRR19/cron_expression_parser/cmd/internal/cronparser"
	"github.com/stretchr/testify/assert"
)

type cronServiceConfig struct {
	t       *testing.T
	service cronparser.CronParser
}

func newCronServiceConfig(t *testing.T) *cronServiceConfig {
	return &cronServiceConfig{
		t:       t,
		service: cronparser.New(),
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		description string
		expression  string
		err         error
		expected    *cronparser.CronFields
	}{
		{
			description: "valid expression",
			expression:  "*/15 0 1,15 * 1-5 /usr/bin/find",
			err:         nil,
			expected: &cronparser.CronFields{
				Minutes:    []int{0, 15, 30, 45},
				Hours:      []int{0},
				DayOfMonth: []int{1, 15},
				Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				DayOfWeek:  []int{1, 2, 3, 4, 5},
				Command:    stringPtr("/usr/bin/find"),
			},
		},
		{
			description: "invalid expression",
			expression:  "invalid expression",
			err:         cronparser.ErrInvalidExpression,
			expected:    nil,
		},
		{
			description: "empty expression",
			expression:  "",
			err:         cronparser.ErrInvalidExpression,
			expected:    nil,
		},
	}

	ctrl := newCronServiceConfig(t)
	for _, test := range tests {
		result, err := ctrl.service.Parse(test.expression)
		assert.Error(t, test.err, err)
		assert.Equal(t, test.expected, result)
	}
}

func stringPtr(s string) *string {
	return &s
}
