package travel

import "errors"

type BookingContext struct {
	strategy CostStrategy
}

func (c *BookingContext) SetStrategy(s CostStrategy) {
	c.strategy = s
}

func (c *BookingContext) Calculate(req TripRequest) (float64, error) {
	if c.strategy == nil {
		return 0, errors.New("strategy not set")
	}
	return c.strategy.Calculate(req)
}