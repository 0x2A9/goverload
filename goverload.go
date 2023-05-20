package goverload

import (
	"lamia-mortis/goverload/requests"
)

func NewOverloader[T goverload.IRequestBodyType]() *Overloader[T] {
	return &Overloader[T]{
		Runners: map[string]*Runner[T]{},
	}
}

func NewRequest[T goverload.IRequestBodyType](kind string, name string) goverload.IRequest[T] {
	// validator.isValid(name)

	switch kind {
	case "http":
		return NewHttpRequest[T](name)
	case "ws": 
	    return NewWsRequest[T](name)	
	default: 
	    return &goverload.Request[T]{
			Name: name,
		}	
	}
}

func NewHttpRequest[T goverload.IRequestBodyType](name string) *goverload.HttpRequest[T] {
	return &goverload.HttpRequest[T]{
		Request: goverload.Request[T]{
			Name: name,
		},
	}
}

func NewWsRequest[T goverload.IRequestBodyType](name string) *goverload.WsRequest[T] {
	return &goverload.WsRequest[T]{
		Request: goverload.Request[T]{
			Name: name,
		},
	}
}
