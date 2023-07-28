package internal

import "fmt"

type DefaultHandler struct{}

func (dh *DefaultHandler) SetNext(handler Handler) {
	// Não é necessário fazer nada, pois este é o último manipulador
}

func (dh *DefaultHandler) HandleRequest(request Request) {
	fmt.Println("Nenhum manipulador disponível para lidar com a requisição")
}
