package goverload 

type WsRequest[T IRequestBodyType] struct {
	Request[T]
}
