package requests

type HttpRequest[RBT IRequestBodyType] struct {
	Request[RBT]
	Method string        
}

func (req *HttpRequest[RBT]) SetMethod(method string) IRequest[RBT] {
	// validator.isValid(method)

	req.Method = method
	return req
}
