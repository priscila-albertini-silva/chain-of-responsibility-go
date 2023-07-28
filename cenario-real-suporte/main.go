package main

import "github.com-pessoal/priscila-albertini/chain-of-responsibility/cenario-real-suporte/internal"

func main() {
	// Criação dos manipuladores
	basicHandler := &internal.BasicSupportHandler{}
	advancedHandler := &internal.AdvancedSupportHandler{}
	expertHandler := &internal.ExpertSupportHandler{}

	// Configuração da cadeia de responsabilidade
	basicHandler.SetNext(advancedHandler)
	advancedHandler.SetNext(expertHandler)

	// Criação das solicitações de suporte
	request1 := internal.SupportRequest{Level: internal.Basic, Message: "Meu computador não liga."}
	request2 := internal.SupportRequest{Level: internal.Advanced, Message: "Não consigo conectar à Internet."}
	request3 := internal.SupportRequest{Level: internal.Expert, Message: "Minha rede local não está funcionando."}
	request4 := internal.SupportRequest{Level: internal.Basic, Message: "Estou com problemas para imprimir."}

	// Envio das solicitações pela cadeia
	basicHandler.HandleRequest(request1)
	basicHandler.HandleRequest(request2)
	basicHandler.HandleRequest(request3)
	basicHandler.HandleRequest(request4)
}
