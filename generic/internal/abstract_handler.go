package internal

// Estrutura base
type AbstractHandler struct {
	nextHandler Handler
}

func (ah *AbstractHandler) SetNext(handler Handler) {
	ah.nextHandler = handler
}

func (ah *AbstractHandler) HandleRequest(request Request) {
	if ah.nextHandler != nil {
		ah.nextHandler.HandleRequest(request)
	}
}
