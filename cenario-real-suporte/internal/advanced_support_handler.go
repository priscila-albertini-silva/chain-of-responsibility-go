package internal

import "fmt"

type AdvancedSupportHandler struct {
	AbstractSupportHandler
}

func (ash *AdvancedSupportHandler) HandleRequest(request SupportRequest) {
	if request.Level == Advanced {
		fmt.Println("Suporte Avançado lida com a solicitação:", request.Message)
	} else {
		ash.AbstractSupportHandler.HandleRequest(request)
	}
}
