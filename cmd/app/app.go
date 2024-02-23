package app

import (
	"github.com/KRR19/cron_expression_parser/cmd/internal/console"
	"github.com/KRR19/cron_expression_parser/cmd/internal/cronparser"
)

func Start() {
	parser := cronparser.New()
	console.New(parser).Run()
}
