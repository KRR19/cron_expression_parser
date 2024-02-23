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

		printCronFields(result)
	}

	fmt.Println("Quitting cron expression parser...")
}

func (s *service) readln() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(scanner.Text())
}

func printCronFields(fields *cronparser.CronFields) {
	fmt.Println()

	printArray("minute", fields.Minutes)
	printArray("hour", fields.Hours)
	printArray("day of month", fields.DayOfMonth)
	printArray("month", fields.Month)
	printArray("day of week", fields.DayOfWeek)

	fmt.Printf("command       %s\n\n", fields.Command)
}

func printArray(name string, values []int) {
	formattedValues := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(values)), " "), "[]")
	fmt.Printf("%-14s %s\n", name, formattedValues)
}
