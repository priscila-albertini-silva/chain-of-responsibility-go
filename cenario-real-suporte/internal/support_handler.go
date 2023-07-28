package internal

type SupportHandler interface {
	SetNext(handler SupportHandler)
	HandleRequest(request SupportRequest)
}
