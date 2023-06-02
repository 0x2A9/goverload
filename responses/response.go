package responses 

type IResponse interface {
	GetBody() map[string]any
}

type Response struct {
	Headers  map[string]string
	Body     map[string]any
}

func (res *Response) GetBody() map[string]any {
	return res.Body
}