package main

import "github.com-pessoal/priscila-albertini/chain-of-responsibility/internal"

func main() {

	// Criação dos manipuladores
	handlerA := &internal.ConcreteHandlerA{}
	handlerB := &internal.ConcreteHandlerB{}
	handlerC := &internal.DefaultHandler{}

	// Configuração da cadeia de responsabilidade
	handlerA.SetNext(handlerB)
	handlerB.SetNext(handlerC)

	// Criação da requisição
	request1 := internal.Request{Type: "A"}
	request2 := internal.Request{Type: "B"}
	request3 := internal.Request{Type: "C"}

	// Envio das requisições pela cadeia
	handlerA.HandleRequest(request1)
	handlerA.HandleRequest(request2)
	handlerA.HandleRequest(request3)
}
