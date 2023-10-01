package goverload

import (
	"lamia-mortis/goverload/requests"
	"lamia-mortis/goverload/responses"
	"lamia-mortis/goverload/stdout"
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
	pb := stdout.NewProgressBar(0, o.GetTotalAmountForAllRunners())
	totalCount := 0

	for reqName, runner := range o.Runners {
		go func(
			reqName string,
			resChan chan responses.IResponse,
			quitChan chan bool,
			errChan chan error,
			runner *Runner[RBT],
			pb *stdout.Bar,
		) {

		CurrentRunner:
			for {
				select {
				case err := <-errChan:
					panic("Error during the %s request execution:" + reqName + "\n" + err.Error())
				case <-resChan:
					totalCount++
					pb.Render(uint64(totalCount))
				case <-quitChan:
					break CurrentRunner
				}
			}
		}(reqName, resChan, quitChan, errChan, runner, pb)

		runner.Run()
	}

	pb.Finish()

	return true
}

func (o *Overloader[RBT]) GetTotalAmountForAllRunners() uint64 {
	var total uint64 = 0

	for _, runner := range o.Runners {
		total += uint64(runner.Config.Amount)
	}

	return total
}
