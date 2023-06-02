package requests

type WsRequest[RBT IRequestBodyType] struct {
	Request[RBT]
}
