package goverload

import (
	"fmt"
	"lamia-mortis/goverload/drivers"
	"lamia-mortis/goverload/helpers/enums"
	"lamia-mortis/goverload/requests"
)

type IHandler[RBT requests.IRequestBodyType] interface {
	Send(requests.IRequest[RBT]) 
}

type Handler[RBT requests.IRequestBodyType] struct {
	Driver drivers.IDriver[RBT]
}

func (h *Handler[RBT]) Send(req requests.IRequest[RBT]) {
	res, err := h.Driver.Send(req)
	
	if err != nil {
		errChan <- fmt.Errorf("error during sending the %s request: %s", req.GetName(), err.Error())
	}

	resChan <- res
}

func NewHandler[RBT requests.IRequestBodyType](reqType string) IHandler[RBT] {
	var driver drivers.IDriver[RBT]

	switch reqType {
	case enums.HTTP.String():
		driver = &drivers.HttpAdapter[RBT]{}
	case enums.WS.String(): 
	    driver = &drivers.WsAdapter[RBT]{}
	default: 
	    panic("Driver for the selected protocol is not exist")
	}

	return &Handler[RBT]{
		Driver: driver,
	}
}
