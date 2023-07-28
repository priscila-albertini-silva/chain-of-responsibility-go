package internal

import "fmt"

type ConcreteHandlerA struct {
	AbstractHandler
}

func (cha *ConcreteHandlerA) HandleRequest(r Request) {
	if r.Type == "A" {
		fmt.Println("Handler A lida com a requisição")
	} else {
		cha.AbstractHandler.HandleRequest(r)
	}
}
