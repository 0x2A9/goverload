package handlers

import (
	"fmt"

	"lamia-mortis/goverload/handlers/drivers"
	"lamia-mortis/goverload/requests"
	"lamia-mortis/goverload/responses"
)

type HttpHandler[RBT requests.IRequestBodyType] struct {
	*Handler[RBT]
}

func (hh *HttpHandler[RBT]) Send(req requests.IRequest[RBT]) responses.IResponse {
	res, err := hh.Driver.Send(req)
	
	if err != nil {
		fmt.Errorf("Error: ", err.Error())
	}

	return res
}

func NewHttpHandler[RBT requests.IRequestBodyType]() IHandler[RBT] {
	return &HttpHandler[RBT]{
		Handler: &Handler[RBT]{
			Driver: &drivers.HttpAdapter[RBT]{},
		},
	}
}