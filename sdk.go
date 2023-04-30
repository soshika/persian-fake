package persian_fake

type PersianFakeServicesInterface interface {
	Sentences(count int8) (*string, error)
	Words(count int16) ([]string, error)
	Address() (*string, error)
}

func Sentences(count int8) (*string, error) {

	return nil, nil
}

func Words(count int16) ([]string, error) {
	return nil, nil
}

func Address() (*string, error) {
	return nil, nil
}
