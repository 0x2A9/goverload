package goverload

import (
	"lamia-mortis/goverload/requests"
	"lamia-mortis/goverload/handlers"
)

type Overloader[RBT requests.IRequestBodyType] struct {
	Runners map[string]*Runner[RBT]
}

func (o *Overloader[RBT]) AddRequest(req requests.IRequest[RBT]) *Runner[RBT] {
	var handler handlers.IHandler[RBT]

	switch req.(type) {
	    case *requests.HttpRequest[RBT]:
		    handler = handlers.NewHttpHandler[RBT]()				
	}

	o.Runners[req.GetName()] = &Runner[RBT]{ 
		Request: req, 
		Handler: handler,
		Config: &RunnerConfig{
			Amount:     0,
			Frequency:  0,
		},
	}

	return o.Runners[req.GetName()]
}
