package goverload

type IRequest[T IRequestBodyType] interface {
	GetName()                      string
	SetOrigin(string)              IRequest[T]
	SetPath(string)                IRequest[T]
	SetParams(map[string]string)   IRequest[T]
	SetHeaders(map[string]string)  IRequest[T]
	SetBody(T)                     IRequest[T]
}

type Request[T IRequestBodyType] struct {
	// upper case for the property name for having possibility to use it from other packages, files and folders 
	// unknown field path in struct literal of type "lamia-mortis/goverload/requests".Request[T]
	Name     string              
	Origin   string
	Path     string 
	Params   map[string]string
	Headers  map[string]string
	Body     T
}

func (r *Request[T]) GetName() string {
	return r.Name
}

func (r *Request[T]) SetOrigin(origin string) IRequest[T] {
	// validator.isValid(origin)

	r.Origin = origin 
	return r
}

func (r *Request[T]) SetPath(path string) IRequest[T] {
	// validator.isValid(path)

	r.Path = path
	return r
}

func (r *Request[T]) SetParams(params map[string]string) IRequest[T] {
	// validator.isValid(params)

	r.Params = params
	return r
}

func (r *Request[T]) SetHeaders(headers map[string]string) IRequest[T] {
	r.Headers = headers
	return r
}

func (r *Request[T]) SetBody(body T) IRequest[T] {
	r.Body = body
	return r
}

type IRequestBodyType interface {
	string | map[string]string
}
