package cronparser

type service struct {

}

func New() *service {
	return &service{}
}

func (s *service) Parse(expression string) (*CronFields, error) {
	return &CronFields{}, nil
}