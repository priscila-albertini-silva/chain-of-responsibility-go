package internal

import "fmt"

type ConcreteHandlerB struct {
	AbstractHandler
}

// Possível solução B
func (chb *ConcreteHandlerB) HandleRequest(request Request) {
	if request.Type == "B" {
		fmt.Println("Handler B lida com a requisição")
	} else {
		chb.AbstractHandler.HandleRequest(request)
	}
}
