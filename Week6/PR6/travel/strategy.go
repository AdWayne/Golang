package travel

type CostStrategy interface {
	Name() string
	Calculate(req TripRequest) (float64, error)
}