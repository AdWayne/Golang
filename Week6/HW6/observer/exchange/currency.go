package exchange

type CurrencyExchange struct {
	observers []IObserver
	rate      float64
}

func (c *CurrencyExchange) Register(observer IObserver) {
	c.observers = append(c.observers, observer)
}

func (c *CurrencyExchange) Remove(observer IObserver) {
	for i, obs := range c.observers {
		if obs == observer {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			break
		}
	}
}

func (c *CurrencyExchange) SetRate(rate float64) {
	c.rate = rate
	c.Notify()
}

func (c *CurrencyExchange) Notify() {
	for _, observer := range c.observers {
		observer.Update(c.rate)
	}
}