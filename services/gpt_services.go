package services

var (
	GPTService gptServiceInterface = &gptService{}
)

type gptService struct{}

type gptServiceInterface interface {
	test() error
}

func (s *gptService) test() error {
	return nil
}
