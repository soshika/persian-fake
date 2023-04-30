package address

const (
	Asia         = 1
	Antarctica   = 2
	NorthAmerica = 3
	Europe       = 4
	Africa       = 5
	SouthAmerica = 6
)

type Address struct {
	Continent string  `json:"continent"`
	Country   string  `json:"country"`
	State     string  `json:"state"`
	City      string  `json:"city"`
	Street    string  `json:"street"`
	Zip       string  `json:"zip"`
	Latitude  string  `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (add *Address) Fake() {

}
