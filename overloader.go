package goverload 

import (
	"lamia-mortis/goverload/requests"
)

type Overloader[T goverload.IRequestBodyType] struct {
	Runners map[string]*Runner[T]
}

func (o *Overloader[T]) AddRequest(r goverload.IRequest[T]) *Runner[T] {
	o.Runners[r.GetName()] = &Runner[T]{ 
		Request: r, 
		Config: &RunnerConfig{
			Amount:     0,
			Frequency:  0,
		},
	}

	return o.Runners[r.GetName()]
}
