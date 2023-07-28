package internal

type AbstractHandler struct {
	nextHandler Handler
}

func (ah *AbstractHandler) SetNext(h Handler) {
	ah.nextHandler = h
}

func (ah *AbstractHandler) HandleRequest(r Request) {
	if ah.nextHandler != nil {
		ah.nextHandler.HandleRequest(r)
	}
}
