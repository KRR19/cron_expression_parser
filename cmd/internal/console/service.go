package console

import "fmt"

type Service struct {

}

func New() *Service {
	return &Service{}
}

func (s *Service) Run() {
	fmt.Println("Running cron expression parser...")

	for {
		fmt.Println("Enter a cron expression in correct format. Example '*/15 0 1,15 * 1-5 /usr/bin/find'. Press 'q' to quit.")
		var input string
		fmt.Scanln(&input)
		if input == quitCommand{
			break
		}
		fmt.Println("You entered:", input)
	}
}