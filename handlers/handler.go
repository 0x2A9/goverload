package handlers

import (
	"fmt"

	"lamia-mortis/goverload/drivers"
	"lamia-mortis/goverload/helpers"
	"lamia-mortis/goverload/requests"
	"lamia-mortis/goverload/responses"
)

type IHandler[RBT requests.IRequestBodyType] interface {
	Send(requests.IRequest[RBT]) responses.IResponse
}

type Handler[RBT requests.IRequestBodyType] struct {
	Driver drivers.IDriver[RBT]
}

func (h *Handler[RBT]) Send(req requests.IRequest[RBT]) responses.IResponse {
	res, err := h.Driver.Send(req)
	
	if err != nil {
		fmt.Errorf("Error: ", err.Error())
	}

	return res
}

func NewHandler[RBT requests.IRequestBodyType](reqType string) IHandler[RBT] {
	var driver drivers.IDriver[RBT]

	switch reqType {
	case helpers.HTTP.String():
		driver = &drivers.HttpAdapter[RBT]{}
	case helpers.WS.String(): 
	    driver = &drivers.WsAdapter[RBT]{}
	default: 
	    panic("Driver for the selected protocol is not exist")
	}

	return &Handler[RBT]{
		Driver: driver,
	}
}