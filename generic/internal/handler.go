package internal

// Manipulador
type Handler interface {
	SetNext(h Handler)
	HandleRequest(r Request)
}
