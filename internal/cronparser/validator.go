package cronparser

import (
	"log"
	"regexp"
)

func (s *service) isValidCronExpression(cronString string) bool {
	cronRegex := `^(\*|[0-9\-\/,\*]+) (\*|[0-9\-\/,\*]+) (\*|[0-9\-\/,\*]+) (\*|[0-9\-\/,\*]+) (\*|[0-9\-\/,\*]+)(\s(\*|[0-9\-\/,\*]+))?(\s([a-zA-Z0-9\/\-_\.\s\*\[\]\(\)\{\}\@]+))?$`
	matched, err := regexp.MatchString(cronRegex, cronString)
	if err != nil {
		log.Fatalf("Error while validating cron expression: %v", err)
	}
	return matched
}
