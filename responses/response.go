package responses

import (
	"encoding/json"
	"fmt"
)

type IResponse interface {
	GetHeaders()        map[string]any
	GetHeadersString()  string
	GetBody()           map[string]any
	GetBodyString()     string
	toString(any)       string
}

type Response struct {
	Headers  map[string]any
	Body     map[string]any
}

func (res *Response) GetHeaders() map[string]any {
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

func NewResponse(headers map[string]any, body []byte) IResponse {
	var content map[string]any

	err := json.Unmarshal(body, &content)

	if err != nil {
		fmt.Printf("Response JSON parsing error: " + err.Error())
	}

	return &Response{
		Headers: headers,
		Body:    content,
	}
}