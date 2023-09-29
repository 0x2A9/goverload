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
		Config: &RunnerConfig{
			Amount:    0,
			Frequency: 0,
		},
	}

	return o.Runners[req.GetName()]
}

func (o *Overloader[RBT]) Run() bool {
	for reqName, runner := range o.Runners {
		runner.Run()

		count := 0

		for {
			select {
			case err := <-errChan:
				fmt.Printf("\n\033[31mError during the %s request execution:", reqName)
				fmt.Printf("\n%s", err.Error())
				return false
			case res := <-resChan:
				count++
				fmt.Printf("\n\033[32mRespone #%d for the %s request:", count, reqName)
				fmt.Printf("\n\n\033[37mHeaders \n%s\n\nBody \n%s\n\n", res.GetHeadersString(), res.GetBodyString())
			case <-quitChan:
				return true
			}
		}
	}

	return true
}
