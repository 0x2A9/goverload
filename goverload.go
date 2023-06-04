package goverload

import (
	"lamia-mortis/goverload/helpers"
	"lamia-mortis/goverload/requests"
)

func NewOverloader[RBT requests.IRequestBodyType]() *Overloader[RBT] {
	return &Overloader[RBT]{
		Runners: map[string]*Runner[RBT]{},
	}
}

func NewRequest[RBT requests.IRequestBodyType](p helpers.Protocol, name string) requests.IRequest[RBT] {
	// validator.isValid(name)

	switch p {
	case helpers.HTTP:
		return NewHttpRequest[RBT](name, p.String())
	case helpers.WS: 
	    return NewWsRequest[RBT](name, p.String())	
	default: 
	    panic("The request protocol is not supported")
	}
}

func NewHttpRequest[RBT requests.IRequestBodyType](name string, protocol string) *requests.HttpRequest[RBT] {
	return &requests.HttpRequest[RBT]{
		Request: requests.Request[RBT]{
			Name:     name,
			Protocol: protocol,
		},
	}
}

func NewWsRequest[RBT requests.IRequestBodyType](name string, protocol string) *requests.WsRequest[RBT] {
	return &requests.WsRequest[RBT]{
		Request: requests.Request[RBT]{
			Name:     name,
			Protocol: protocol,
		},
	}
}
