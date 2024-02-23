package cronparser

type service struct {
}

func New() *service {
	return &service{}
}

func (s *service) Parse(expression string) (*CronFields, error) {
	if !s.isValidCronExpression(expression) {
		return nil, ErrInvalidExpression
	}
	return &CronFields{}, nil
}
