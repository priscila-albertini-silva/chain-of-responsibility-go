package internal

type AbstractSupportHandler struct {
	nextHandler SupportHandler
}

func (ah *AbstractSupportHandler) SetNext(handler SupportHandler) {
	ah.nextHandler = handler
}

func (ah *AbstractSupportHandler) HandleRequest(request SupportRequest) {
	if ah.nextHandler != nil {
		ah.nextHandler.HandleRequest(request)
	}
}
