package goverload

import (
	"sync"
	"time"
	"lamia-mortis/goverload/requests"
	"lamia-mortis/goverload/responses"
)

var resChan  chan responses.IResponse = make(chan responses.IResponse)
var errChan  chan error = make(chan error)
var quitChan chan bool  = make(chan bool) 

type Runner[RBT requests.IRequestBodyType] struct {
	Request requests.IRequest[RBT]
	Handler IHandler[RBT]
	Config  *RunnerConfig
}

func (r *Runner[RBT]) SetConfig(amount uint16, frequency uint16) *Runner[RBT] {
	r.Config.Amount     = amount
	r.Config.Frequency  = frequency

	return r
}

func (r *Runner[RBT]) Run() {
	interval := r.calcInterval()
	wg       := &sync.WaitGroup{}

	for i := 0; i < int(r.Config.Amount); i++ {
		wg.Add(1)

		go func(req requests.IRequest[RBT]) {
			defer wg.Done()
			r.Handler.Send(req)
		}(r.Request)

		time.Sleep(time.Duration(interval) * time.Nanosecond)
	}

	// by default a chan holds no items, so all goroutines are blocked on sending, until something reads from it 
	// otherwise goroutines never reach the wg.Done() statement 
	// so closing the channel in it's own goroutine (while reading is executed from the main thread)
	go func() {
		// wait for all goroutines to complete
		wg.Wait() 

		quitChan <- true

		// closing channels after completion of wait for goroutines
		close(errChan) 
		close(resChan)
	}()
}

/**
 *  return interval (per 1 second) between requests in nanoseconds
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
