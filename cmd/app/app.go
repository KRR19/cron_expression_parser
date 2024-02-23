package app

import (
	"github.com/KRR19/cron_expression_parser/internal/console"
	"github.com/KRR19/cron_expression_parser/internal/cronparser"
)

func Start() {
	parser := cronparser.New()
	console.New(parser).Run()
}
