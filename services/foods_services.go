package services

var (
	FoodsService foodsServiceInterface = &foodsService{}
)

type foodsService struct{}

type foodsServiceInterface interface {
	test() error
}

func (s *foodsService) test() error {
	return nil
}
