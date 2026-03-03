package exchange

type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify()
}