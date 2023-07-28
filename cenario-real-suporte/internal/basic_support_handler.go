package internal

import "fmt"

type BasicSupportHandler struct {
	AbstractSupportHandler
}

func (bsh *BasicSupportHandler) HandleRequest(request SupportRequest) {
	if request.Level == Basic {
		fmt.Println("Suporte Básico lida com a solicitação:", request.Message)
	} else {
		bsh.AbstractSupportHandler.HandleRequest(request)
	}
}
