package drivers

import (
	"github.com/lamia-mortis/goverload/requests"
	"github.com/lamia-mortis/goverload/responses"
)

type IDriver[RBT requests.IRequestBodyType] interface {
	Send(requests.IRequest[RBT]) (responses.IResponse, error)
}