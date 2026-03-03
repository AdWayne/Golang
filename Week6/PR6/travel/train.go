package travel

import "math"

type TrainStrategy struct{}

func (t TrainStrategy) Name() string {
	return "Train"
}

func (t TrainStrategy) Calculate(req TripRequest) (float64, error) {
	if err := req.Validate(); err != nil {
		return 0, err
	}

	base := req.DistanceKm * 0.1

	if req.Class == Business {
		base *= 1.4
	}

	base *= float64(req.Passengers)

	base = base * (1 - req.DiscountPct/100)

	return math.Round(base*100) / 100, nil
}