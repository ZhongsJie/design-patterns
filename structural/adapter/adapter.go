package main

type Handler interface {
	Handle(in int) int
}

type HandlerAdapter struct {
	Handler Handler
}

func NewHandlerAdapter(handler Handler) HandlerAdapter {
	return HandlerAdapter{Handler: handler}
}

func (h HandlerAdapter) Handle(in int) int {
	return h.Handler.Handle(in)
}
