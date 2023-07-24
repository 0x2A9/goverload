package requests

type IRequest[RBT IRequestBodyType] interface {
	GetName()                      string
	GetOrigin()                    string
	GetProtocol()                  string
	SetHost(string)                IRequest[RBT]
	SetPort(string)                IRequest[RBT]
	SetPath(string)                IRequest[RBT]
	SetParams(map[string]string)   IRequest[RBT]
	SetHeaders(map[string]string)  IRequest[RBT]
	SetBody(RBT)                   IRequest[RBT]
	Type()                         string
	Send()                         
}

type Request[RBT IRequestBodyType] struct {
	Name     string      
	Protocol string
	Host     string 
	Port     string
	Path     string 
	Params   map[string]string
	Headers  map[string]string
	Body     RBT
}

func (req *Request[RBT]) GetName() string {
	return req.Name
}

func (req *Request[RBT]) GetOrigin() string {
	return req.Protocol + "://" + req.Host + ":" + req.Port
}

func (req *Request[RBT]) GetProtocol() string {
	return req.Protocol
}

func (req *Request[RBT]) SetHost(host string) IRequest[RBT] {
	// validator.isValid(host)

	req.Host = host
	return req
}

func (req *Request[RBT]) SetPort(port string) IRequest[RBT] {
	// validator.isValid(port)

	req.Port = port
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

func (req *Request[RBT]) Type() string {
	return req.GetProtocol()
}

func (req *Request[RBT]) Send() {

}

type IRequestBodyType interface {
	string | []byte
}
