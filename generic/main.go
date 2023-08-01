package main

import "github.com/priscila-albertini/chain-of-responsibility/generic/internal"

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
	handlerA.HandleRequest(request1) // Saída: "Handler A lida com a requisição"
	handlerA.HandleRequest(request2) // Saída: "Handler B lida com a requisição"
	handlerA.HandleRequest(request3) // Saída: "Nenhum manipulador disponível para lidar com a requisição"
}
