package goverload

import (
	"sync"
	"time"
	"lamia-mortis/goverload/requests"
	"lamia-mortis/goverload/responses"
)

var resChan chan responses.IResponse  = make(chan responses.IResponse, 1)
var errChan chan error                = make(chan error, 1)

type Runner[RBT requests.IRequestBodyType] struct {
	Request  requests.IRequest[RBT]
	Handler  IHandler[RBT]
	Config   *RunnerConfig
}

func (r *Runner[RBT]) SetConfig(amount uint16, frequency uint16) *Runner[RBT] {
	r.Config.Amount     = amount
	r.Config.Frequency  = frequency

	return r
}

func (r *Runner[RBT]) Run() error {
	interval := r.calcInterval()
	var wg sync.WaitGroup

    for i := 0; i < int(r.Config.Amount); i++ {
		wg.Add(1)
		go r.Handler.Send(r.Request)

		err := <- errChan
		if err != nil {
			return err
		}

		time.Sleep(time.Duration(interval) * time.Nanosecond)
    }
	
	wg.Wait()
	return nil
}

/**
 *  return interval between requests in nanoseconds
 */
func (r *Runner[RBT]) calcInterval() uint32 {
	var second uint32 = 1000000000
	return second / uint32(r.Config.Frequency)
}

type RunnerConfig struct {
	/* RPS */
	Frequency  uint16  
	Amount     uint16
}
