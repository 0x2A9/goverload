package goverload

import (
	"fmt"
	"lamia-mortis/goverload/requests"
)

type Overloader[RBT requests.IRequestBodyType] struct {
	Runners map[string]*Runner[RBT]
}

func (o *Overloader[RBT]) AddRequest(req requests.IRequest[RBT]) *Runner[RBT] {
	o.Runners[req.GetName()] = &Runner[RBT]{ 
		Request: req, 
		Handler: NewHandler[RBT](req.Type()),
		Config:  &RunnerConfig{
			Amount:     0,
			Frequency:  0,
		},
	}

	return o.Runners[req.GetName()]
}

func (o *Overloader[RBT]) Run() bool {
	for reqName, runner := range o.Runners {
		err := runner.Run()

		if err != nil {
			fmt.Printf("Error during the %s request execution: %s", reqName, err.Error())
			return false
		}
	} 

	return true
}
