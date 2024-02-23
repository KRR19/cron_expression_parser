package cronparser

import (
	"strings"

	"github.com/KRR19/cron_expression_parser/internal/pkg/utils"
)

func (s *service) parse(expression string) (*CronFields, error) {
	fields := strings.Fields(expression)
	
	minutes := s.parseField(fields[0], 0, 59)
	hours := s.parseField(fields[1], 0, 23)
	dayOfMonth := s.parseField(fields[2], 1, 31)
	month := s.parseField(fields[3], 1, 12)
	dayOfWeek := s.parseField(fields[4], 0, 6)

	return &CronFields{
		Minutes:    minutes,
		Hours:      hours,
		DayOfMonth: dayOfMonth,
		Month:      month,
		DayOfWeek:  dayOfWeek,
		Command:    utils.StringPtr(fields[5]),
	}, nil
}
