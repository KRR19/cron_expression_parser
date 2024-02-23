package cronparser

import (
	"log"
	"strconv"
	"strings"

	"github.com/KRR19/cron_expression_parser/internal/pkg/utils"
)

func (s *service) parse(expression string) (*CronFields, error) {
	fields := strings.Fields(expression)

	minutes := s.parseField(fields[0], minMinutes, maxMinutes)
	hours := s.parseField(fields[1], minHours, maxHours)
	dayOfMonth := s.parseField(fields[2], minDay, maxDay)
	month := s.parseField(fields[3], minMonth, maxMonth)
	dayOfWeek := s.parseField(fields[4], minWeek, maxWeek)

	return &CronFields{
		Minutes:    minutes,
		Hours:      hours,
		DayOfMonth: dayOfMonth,
		Month:      month,
		DayOfWeek:  dayOfWeek,
		Command:    utils.StringPtr(strings.Join(fields[5:], " ")),
	}, nil
}

func (s *service) parseField(field string, min, max int) []int {
	if field == "*" {
		return s.allValues(min, max)
	} else if strings.Contains(field, "/") {
		return s.parseStep(field, max)
	} else if strings.Contains(field, "-") {
		return s.parseRange(field, min, max)
	} else if strings.Contains(field, ",") {
		return s.parseComma(field, min, max)
	}

	return s.parseSingle(field, min, max)
}

func (s *service) parseSingle(field string, min, max int) []int {
	value, err := strconv.Atoi(field)
	if err != nil {
		log.Printf("Error parsing single value: %s", err)
		return nil
	}

	if value >= min && value <= max {
		return []int{value}
	}

	return nil
}

func (s *service) parseComma(field string, min, max int) []int {
	result := make([]int, 0)
	parts := strings.Split(field, ",")
	for _, part := range parts {
		value, err := strconv.Atoi(part)
		if err != nil {
			log.Printf("Error parsing comma-separated value: %s", err)
			continue
		}
		if value >= min && value <= max {
			result = append(result, value)
		}
	}
	return result
}

func (s *service) parseRange(field string, min, max int) []int {
	rangeParts := strings.Split(field, "-")
	start, err := strconv.Atoi(rangeParts[0])
	if err != nil {
		log.Printf("Error parsing range start: %s", err)
		return nil
	}

	end, err := strconv.Atoi(rangeParts[1])
	if err != nil {
		log.Printf("Error parsing range end: %s", err)
		return nil
	}

	result := make([]int, 0, max-min+1)
	for i := start; i <= end; i++ {
		result = append(result, i)
	}

	return result
}

func (s *service) parseStep(field string, max int) []int {
	stepParts := strings.Split(field, "/")
	step, err := strconv.Atoi(stepParts[1])
	if err != nil {
		log.Printf("Error parsing step: %s", err)
		return nil
	}

	result := make([]int, 0)
	for i := 0; i <= max; i += step {
		result = append(result, i)
	}

	return result
}

func (s *service) allValues(min, max int) []int {
	values := make([]int, max-min+1)
	for i := min; i <= max; i++ {
		values[i-min] = i
	}
	return values
}
