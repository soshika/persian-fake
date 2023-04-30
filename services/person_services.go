package services

var (
	PersonService personServiceInterface = &personService{}
)

type personService struct{}

type personServiceInterface interface {
	test() error
}

func (s *personService) test() error {
	return nil
}
