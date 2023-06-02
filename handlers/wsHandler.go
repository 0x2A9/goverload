package handlers

import (
	"fmt"

	"lamia-mortis/goverload/handlers/drivers"
	"lamia-mortis/goverload/requests"
	"lamia-mortis/goverload/responses"
)

type WsHandler[RBT requests.IRequestBodyType] struct {
	*Handler[RBT]
}

func (wh *WsHandler[RBT]) Send(req requests.IRequest[RBT]) responses.IResponse {
	res, err := wh.Driver.Send(req)
	
	if err != nil {
		fmt.Errorf("Error: ", err.Error())
	}

	return res
}

func NewWsHandler[RBT requests.IRequestBodyType]() IHandler[RBT] {
	return &WsHandler[RBT]{
		Handler: &Handler[RBT]{
			Driver: &drivers.WsAdapter[RBT]{},
		},
	}
}