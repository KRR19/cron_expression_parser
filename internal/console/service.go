package console

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/KRR19/cron_expression_parser/internal/cronparser"
)

type service struct {
	parses cronparser.CronParser
}

func New(parses cronparser.CronParser) *service {
	return &service{parses}
}

func (s *service) Run() {
	fmt.Println("Running cron expression parser...")

	s.readFromArgs()

	s.readFromConsole()

	fmt.Println("Quitting cron expression parser...")
}

func (s *service) readFromArgs() {
	if len(os.Args) < 2 {
		return
	}

	result, err := s.parses.Parse(os.Args[1])
	if err != nil {
		if err == cronparser.ErrInvalidExpression {
			fmt.Println("Invalid expression. Please try again.")
		}
		return
	}

	s.printCronFields(result)
}

func (s *service) readFromConsole() {
	for {
		fmt.Println("Enter a cron expression in correct format.\nExample '*/15 0 1,15 * 1-5 /usr/bin/find'.\nPress 'q' to quit.")
		input := s.readln()

		if input == quitCommand {
			break
		}

		result, err := s.parses.Parse(input)
		if err != nil {
			if err == cronparser.ErrInvalidExpression {
				fmt.Println("Invalid expression. Please try again.")
			}
			continue
		}

		s.printCronFields(result)
	}
}

func (s *service) readln() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimFunc(strings.TrimSpace(scanner.Text()), func(r rune) bool {
		return r == '"' || r == '\''
	})
}

func (s *service) printCronFields(fields *cronparser.CronFields) {
	fmt.Println()

	s.printArray("minute", fields.Minutes)
	s.printArray("hour", fields.Hours)
	s.printArray("day of month", fields.DayOfMonth)
	s.printArray("month", fields.Month)
	s.printArray("day of week", fields.DayOfWeek)

	fmt.Printf("command       %s\n\n", fields.Command)
}

func (s *service) printArray(name string, values []int) {
	formattedValues := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(values)), " "), "[]")
	fmt.Printf("%-14s %s\n", name, formattedValues)
}
