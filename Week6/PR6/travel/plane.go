package travel

import "math"

type PlaneStrategy struct{}

func (p PlaneStrategy) Name() string {
	return "Plane"
}

func (p PlaneStrategy) Calculate(req TripRequest) (float64, error) {
	if err := req.Validate(); err != nil {
		return 0, err
	}

	base := req.DistanceKm * 0.2
	if req.Class == Business {
		base *= 1.8
	}

	if req.HasBaggage {
		base += 20
	}

	base *= float64(req.Passengers)

	base = base * (1 - req.DiscountPct/100)

	return math.Round(base*100) / 100, nil
}