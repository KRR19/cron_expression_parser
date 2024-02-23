package console

import (
	"fmt"
	"strings"

	"github.com/KRR19/cron_expression_parser/cmd/internal/cronparser"
)

type Service struct {
	parses cronparser.CronParser
}

func New(parses cronparser.CronParser) *Service {
	return &Service{parses}
}

func (s *Service) Run() {
	fmt.Println("Running cron expression parser...")

	for {
		fmt.Println("Enter a cron expression in correct format. Example '*/15 0 1,15 * 1-5 /usr/bin/find'. Press 'q' to quit.")
		var input string
		fmt.Scanln(&input)
		input = strings.TrimSpace(input)
		if input == quitCommand{
			break
		}

	}
}