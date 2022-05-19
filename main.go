package main

import (
	"fmt"
	"obs/biz"
	"time"
)

func backgroundFlow(observer biz.IObserver, done chan<- string) {
	var observable biz.IObservable = biz.NewSubject("Subj01")
	observable.Subscribe(observer)
	observable.ListSubscribers()
	observable.OnChange(biz.Status{Msg: "Executing Activity 1"})
	time.Sleep(1000 * time.Millisecond)
	observable.OnChange(biz.Status{Msg: "Executing Activity 2"})
	time.Sleep(1000 * time.Millisecond)
	observable.OnChange(biz.Status{Msg: "ExecutingActivity 3"})
	time.Sleep(1000 * time.Millisecond)
	observable.Unsubscribe(observer)
	done <- "Y"
}

func mainFlow() {
	var observer biz.IObserver = biz.Listener{Name: "Listener01"}
	done := make(chan string, 1)
	go backgroundFlow(observer, done)
	<-done
}

func main() {
	fmt.Printf("Welcome to Observer \n")
	mainFlow()
}
