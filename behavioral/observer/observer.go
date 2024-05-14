package main

import (
	"fmt"
	"time"
)

type Event struct {
	Data int64
}

type Observer interface {
	OnNotify(event Event)
}

type Notifier interface {
	Register(observer Observer)
	UnRegister(observer Observer)
	Notify(event Event)
}

type (
	eventObserver struct {
		id int
	}

	eventNotifier struct {
		observers map[Observer]struct{}
	}
)

func (e eventNotifier) Register(observer Observer) {
	e.observers[observer] = struct{}{}
}

func (e eventNotifier) UnRegister(observer Observer) {
	delete(e.observers, observer)
}

func (e eventNotifier) Notify(event Event) {
	for o := range e.observers {
		o.OnNotify(event)
	}
}

func (e eventObserver) OnNotify(event Event) {
	fmt.Printf("*** Observer %d received: %d\n", e.id, event.Data)
}

var _ Observer = new(eventObserver)
var _ Notifier = new(eventNotifier)

func main() {
	observerA := eventObserver{id: 1}
	observerB := eventObserver{id: 2}
	notifier := eventNotifier{observers: make(map[Observer]struct{})}
	notifier.Register(observerA)
	notifier.Register(observerB)

	stop := time.NewTimer(5 * time.Second).C
	tick := time.NewTicker(time.Second).C
	for {
		select {
		case <-stop:
			return
		case t := <-tick:
			notifier.Notify(Event{Data: t.UnixNano()})
		}
	}
}
