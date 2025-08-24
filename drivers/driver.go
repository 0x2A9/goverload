package drivers

import (
	"github.com/0x2A9/goverload/requests"
	"github.com/0x2A9/goverload/responses"
)

type IDriver[RBT requests.IRequestBodyType] interface {
	Send(requests.IRequest[RBT]) (responses.IResponse, error)
}