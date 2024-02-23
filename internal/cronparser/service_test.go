package cronparser_test

import (
	"testing"

	"github.com/KRR19/cron_expression_parser/internal/cronparser"
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
	ctrl := newCronServiceConfig(t)
	t.Run("Valid expression", func(t *testing.T) {
		expression := "*/15 0 1,15 * 1-5 /usr/bin/find"
		expected := &cronparser.CronFields{
			Minutes:    []int{0, 15, 30, 45},
			Hours:      []int{0},
			DayOfMonth: []int{1, 15},
			Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			DayOfWeek:  []int{1, 2, 3, 4, 5},
			Command:    "/usr/bin/find",
		}

		actual, err := ctrl.service.Parse(expression)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("Partial written expression", func(t *testing.T) {
		expression := "*/15 0 1,15 *"
		actual, err := ctrl.service.Parse(expression)
		assert.Error(t, cronparser.ErrInvalidExpression, err)
		assert.Nil(t, actual)
	})

	t.Run("Expression with to long command", func(t *testing.T) {
		expression := "*/15 0 1,15 * 1-5 /usr/bin/find extended command"
		expected := &cronparser.CronFields{
			Minutes:    []int{0, 15, 30, 45},
			Hours:      []int{0},
			DayOfMonth: []int{1, 15},
			Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			DayOfWeek:  []int{1, 2, 3, 4, 5},
			Command:    "/usr/bin/find extended command",
		}

		actual, err := ctrl.service.Parse(expression)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Expression without command", func(t *testing.T) {
		expression := "*/15 0 1,15 * 1-5"
		expected := &cronparser.CronFields{
			Minutes:    []int{0, 15, 30, 45},
			Hours:      []int{0},
			DayOfMonth: []int{1, 15},
			Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			DayOfWeek:  []int{1, 2, 3, 4, 5},
		}

		actual, err := ctrl.service.Parse(expression)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("InvalidExpression", func(t *testing.T) {
		expression := "*/f 0 1,k5 * 1-5 /usr/bin/find"

		actual, err := ctrl.service.Parse(expression)

		assert.Error(t, cronparser.ErrInvalidExpression, err)
		assert.Nil(t, actual)
	})
	t.Run("EmptyExpression", func(t *testing.T) {
		expression := ""

		actual, err := ctrl.service.Parse(expression)

		assert.Error(t, cronparser.ErrInvalidExpression, err)
		assert.Nil(t, actual)
	})
}
