package biz

type Status struct {
	Msg string
}

type IObserver interface {
	GetName() string
	OnChange(status Status)
}

type IObservable interface {
	ListSubscribers() []IObserver
	IndexOfObserver(observer IObserver) int
	Subscribe(observer IObserver)
	Unsubscribe(observer IObserver)
	OnChange(status Status)
}
