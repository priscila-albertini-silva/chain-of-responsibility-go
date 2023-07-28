package internal

import "fmt"

type ExpertSupportHandler struct {
	AbstractSupportHandler
}

func (ash *ExpertSupportHandler) HandleRequest(request SupportRequest) {
	if request.Level == Advanced {
		fmt.Println("Suporte Expert lida com a solicitação:", request.Message)
	} else {
		ash.AbstractSupportHandler.HandleRequest(request)
	}
}
