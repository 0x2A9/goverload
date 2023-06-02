package goverload

import "lamia-mortis/goverload/requests"

func NewOverloader[RBT requests.IRequestBodyType]() *Overloader[RBT] {
	return &Overloader[RBT]{
		Runners: map[string]*Runner[RBT]{},
	}
}

func NewRequest[RBT requests.IRequestBodyType](kind string, name string) requests.IRequest[RBT] {
	// validator.isValid(name)

	switch kind {
	case "http":
		return NewHttpRequest[RBT](name)
	case "ws": 
	    return NewWsRequest[RBT](name)	
	default: 
	    return &requests.Request[RBT]{
			Name: name,
		}	
	}
}

func NewHttpRequest[RBT requests.IRequestBodyType](name string) *requests.HttpRequest[RBT] {
	return &requests.HttpRequest[RBT]{
		Request: requests.Request[RBT]{
			Name: name,
		},
	}
}

func NewWsRequest[RBT requests.IRequestBodyType](name string) *requests.WsRequest[RBT] {
	return &requests.WsRequest[RBT]{
		Request: requests.Request[RBT]{
			Name: name,
		},
	}
}
