package responses 

import (
	"encoding/json"
)

type IResponse interface {
	GetHeaders()        map[string]string
	GetHeadersString()  string
	GetBody()           map[string]any
	GetBodyString()     string
	toString(any)       string
}

type Response struct {
	Headers  map[string]string
	Body     map[string]any
}

func (res *Response) GetHeaders() map[string]string {
	return res.Headers
}

func (res *Response) GetHeadersString() string {
	return res.toString(res.Headers)
}

func (res *Response) GetBody() map[string]any {
	return res.Body
}

func (res *Response) GetBodyString() string {
	return res.toString(res.Body)
}

func (res *Response) toString(value any) string {
	json, err := json.Marshal(value)

	if err != nil {
		panic(err.Error())
	}

	return string(json)
}