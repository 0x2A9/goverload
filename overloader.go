package goverload

import (
	"lamia-mortis/goverload/handlers"
	"lamia-mortis/goverload/requests"
)

type Overloader[RBT requests.IRequestBodyType] struct {
	Runners map[string]*Runner[RBT]
}

func (o *Overloader[RBT]) AddRequest(req requests.IRequest[RBT]) *Runner[RBT] {
	o.Runners[req.GetName()] = &Runner[RBT]{ 
		Request: req, 
		Handler: handlers.NewHandler[RBT](req.Type()),
		Config:  &RunnerConfig{
			Amount:     0,
			Frequency:  0,
		},
	}

	return o.Runners[req.GetName()]
}
