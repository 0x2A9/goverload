package handlers

import (
	"lamia-mortis/goverload/requests"
	"lamia-mortis/goverload/responses"
	"lamia-mortis/goverload/handlers/drivers"
)

type IHandler[RBT requests.IRequestBodyType] interface {
	Send(requests.IRequest[RBT]) responses.IResponse
}

type Handler[RBT requests.IRequestBodyType] struct {
	Driver drivers.IDriver[RBT]
}