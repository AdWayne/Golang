package order

type Item struct {
	Name  string
	Price float64
	Qty   int
}

type Discount interface {
	Apply(amount float64) float64
}

type Order struct {
	ID        int
	Items     []Item
	Discounts []Discount
}

func (o *Order) AddItem(item Item) {
	o.Items = append(o.Items, item)
}

func (o *Order) CalculateTotal() float64 {
	total := 0.0

	for _, item := range o.Items {
		total += item.Price * float64(item.Qty)
	}

	for _, d := range o.Discounts {
		total = d.Apply(total)
	}

	return total
}
