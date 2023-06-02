package requests

type IRequest[RBT IRequestBodyType] interface {
	GetName()                      string
	SetOrigin(string)              IRequest[RBT]
	SetPath(string)                IRequest[RBT]
	SetParams(map[string]string)   IRequest[RBT]
	SetHeaders(map[string]string)  IRequest[RBT]
	SetBody(RBT)                   IRequest[RBT]
	Send()                         
}

type Request[RBT IRequestBodyType] struct {
	Name     string      
	Origin   string
	Path     string 
	Params   map[string]string
	Headers  map[string]string
	Body     RBT
}

func (req *Request[RBT]) GetName() string {
	return req.Name
}

func (req *Request[RBT]) SetOrigin(origin string) IRequest[RBT] {
	// validator.isValid(origin)

	req.Origin = origin 
	return req
}

func (req *Request[RBT]) SetPath(path string) IRequest[RBT] {
	// validator.isValid(path)

	req.Path = path
	return req
}

func (req *Request[RBT]) SetParams(params map[string]string) IRequest[RBT] {
	// validator.isValid(params)

	req.Params = params
	return req
}

func (req *Request[RBT]) SetHeaders(headers map[string]string) IRequest[RBT] {
	req.Headers = headers
	return req
}

func (req *Request[RBT]) SetBody(body RBT) IRequest[RBT] {
	req.Body = body
	return req
}

func (req *Request[RBT]) Send() {

}

type IRequestBodyType interface {
	string | map[string]string
}
