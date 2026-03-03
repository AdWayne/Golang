package travel

import "errors"

type ServiceClass int

const (
	Economy ServiceClass = iota
	Business
)

type PassengerType int

const (
	Adult PassengerType = iota
	Child
	Senior
)

type TripRequest struct {
	DistanceKm    float64
	Class         ServiceClass
	Passengers    int
	PassengerType PassengerType
	HasBaggage    bool
	DiscountPct   float64
}

func (r TripRequest) Validate() error {
	if r.DistanceKm <= 0 {
		return errors.New("distance must be > 0")
	}
	if r.Passengers <= 0 {
		return errors.New("passengers must be > 0")
	}
	if r.DiscountPct < 0 || r.DiscountPct > 100 {
		return errors.New("discount must be 0..100")
	}
	return nil
}