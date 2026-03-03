package travel

import "math"

type BusStrategy struct{}

func (b BusStrategy) Name() string {
	return "Bus"
}

func (b BusStrategy) Calculate(req TripRequest) (float64, error) {
	if err := req.Validate(); err != nil {
		return 0, err
	}

	base := req.DistanceKm * 0.07
	base *= float64(req.Passengers)

	base = base * (1 - req.DiscountPct/100)

	return math.Round(base*100) / 100, nil
}