package goverload

import (
	"lamia-mortis/goverload/requests"
)

type Runner[T goverload.IRequestBodyType] struct {
	Request  goverload.IRequest[T]
	Config   *RunnerConfig
}

func (r *Runner[T]) SetConfig(amount int16, frequency int16) *Runner[T] {
	r.Config.Amount     = amount
	r.Config.Frequency  = frequency

	return r
}

type RunnerConfig struct {
	Amount     int16
	Frequency  int16
}
