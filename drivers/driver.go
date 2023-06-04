package drivers

import (
	"lamia-mortis/goverload/requests"
	"lamia-mortis/goverload/responses"
)

type IDriver[RBT requests.IRequestBodyType] interface {
	Send(requests.IRequest[RBT]) (responses.IResponse, error)
}