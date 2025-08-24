package goverload

import (
	"sync"
	"time"
	"github.com/0x2A9/goverload/requests"
	"github.com/0x2A9/goverload/responses"
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
	}()
}

/**
 *  return interval (per 1 second) between requests in nanoseconds
 */
func (r *Runner[RBT]) calcInterval() int {
	var second int = 1000000000
	var frequency int = int(r.Config.Frequency)

	if frequency > second {
		frequency = second
	}

	return second / frequency
}

type RunnerConfig struct {
	/* RPS can't be more than 1_000_000_000 */
	Frequency  uint16  
	/* Total */
	Amount     uint16
}
