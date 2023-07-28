package internal

import "fmt"

type ConcreteHandlerA struct {
	AbstractHandler
}

// Possível solução A
func (cha *ConcreteHandlerA) HandleRequest(request Request) {
	if request.Type == "A" {
		fmt.Println("Handler A lida com a requisição")
	} else {
		cha.AbstractHandler.HandleRequest(request)
	}
}
