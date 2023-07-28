package internal

import "fmt"

type ConcreteHandlerB struct {
	AbstractHandler
}

func (chb *ConcreteHandlerB) HandleRequest(r Request) {
	if r.Type == "B" {
		fmt.Println("Handler B lida com a requisição")
	} else {
		chb.AbstractHandler.HandleRequest(r)
	}
}
