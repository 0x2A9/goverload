package goverload 

type HttpRequest[T IRequestBodyType] struct {
	Request[T]
}
