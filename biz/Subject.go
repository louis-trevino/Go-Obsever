package biz

import "fmt"

/* Subject is an implementation of IObservable interface */
type Subject struct {
	Name      string
	Listeners []Listener
}

func NewSubject(name string) *Subject {
	listeners := make([]Listener, 0)
	subj := Subject{Name: name, Listeners: listeners}
	return &subj
}

func (subject *Subject) OnChange(status Status) {
	fmt.Printf("Subject: %v Status: %v \n", subject, status)
}

func (subject *Subject) ListSubscribers() []IObserver {
	var observers []IObserver = make([]IObserver, len(subject.Listeners))
	for idx, listener := range subject.Listeners {
		observers[idx] = listener
	}
	return observers
}

func (subject *Subject) IndexOfObserver(observer IObserver) int {
	for idx, lis := range subject.Listeners {
		if lis == observer {
			return idx
		}
	}
	return -1
}

func (subject *Subject) ObserverExists(observer IObserver) bool {
	return subject.IndexOfObserver(observer) >= 0
}

func (subject *Subject) Subscribe(observer IObserver) {
	if subject.ObserverExists(observer) {
		fmt.Printf("Observer %v exists. \n", observer.GetName())
		return
	}
	var listener Listener
	switch observer.(type) {
	case Listener:
		listener = observer.(Listener)
	}
	subject.Listeners = append(subject.Listeners, listener)
	fmt.Printf("Observer %v subscribed.\n", observer.GetName())
}

func RemoveFromListenerSlice(s []Listener, index int) []Listener {
	return append(s[:index], s[index+1:]...)
}

func (subject *Subject) Unsubscribe(observer IObserver) {
	idx := subject.IndexOfObserver(observer)
	if idx < 0 {
		fmt.Printf("Observer %v doesn't exist. \n", observer.GetName())
		return
	}
	subject.Listeners = RemoveFromListenerSlice(subject.Listeners, idx)
	fmt.Printf("Observer %v unsubscribed.\n", observer.GetName())
}
