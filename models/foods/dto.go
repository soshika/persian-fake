package foods

const (
	Fruit     = 0
	Vegetable = 1
	Breakfast = 2
	Lunch     = 3
	Dinner    = 4
	Snack     = 5
	Dessert   = 6
)

type Food struct {
	Name   string `json:"name"`
	Origin string `json:"origin"`
	Type   string `json:"type"`
}

func (food *Food) Fake() {

}

func (food *Food) FakeByOrigin() {

}

func (food *Food) FakeByType() {

}
