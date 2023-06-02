package goverload

import (
	"lamia-mortis/goverload/requests"
	"lamia-mortis/goverload/handlers"
)

type Runner[RBT requests.IRequestBodyType] struct {
	Request  requests.IRequest[RBT]
	Handler  handlers.IHandler[RBT]
	Config   *RunnerConfig
}

func (r *Runner[RBT]) SetConfig(amount int16, frequency int16) *Runner[RBT] {
	r.Config.Amount     = amount
	r.Config.Frequency  = frequency

	return r
}

type RunnerConfig struct {
	Amount     int16
	Frequency  int16
}
