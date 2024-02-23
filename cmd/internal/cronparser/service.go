package cronparser

import (
	"errors"
	"strings"
)

type service struct {
}

func New() *service {
	return &service{}
}

func (s *service) Parse(expression string) (*CronFields, error) {
	if !s.isValidCronExpression(expression) {
		return nil, ErrInvalidExpression
	}

	fields := strings.Fields(expression)
	if len(fields) != 7 {
		return nil, errors.New("invalid cron string format")
	}
	return &CronFields{}, nil
}
