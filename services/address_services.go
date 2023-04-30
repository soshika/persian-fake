package services

var (
	AddressesService addressServiceInterface = &addressesService{}
)

type addressesService struct{}

type addressServiceInterface interface {
	test() error
}

func (s *addressesService) test() error {
	return nil
}
