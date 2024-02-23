package cronparser

type CronParser interface {
	Parse(expression string) (*CronFields, error)
}