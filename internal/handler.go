package internal

type Handler interface {
	SetNext(h Handler)
	HandleRequest(r Request)
}

type Request struct {
	Type string
}
