package biz

import "fmt"

/* Listener is an implementation of IObserver interface */
type Listener struct {
	Name string
}

func (listener Listener) OnChange(status Status) {
	fmt.Printf("Listener %v OnChange. Status: %v \n", listener.GetName(), status)
}

func (listener Listener) GetName() string {
	return listener.Name
}
