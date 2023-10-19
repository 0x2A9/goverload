package drivers

import (
	"bytes"
	"errors"
	"net/http"
	"reflect"

	"github.com/lamia-mortis/goverload/requests"
	"github.com/lamia-mortis/goverload/responses"
)

type HttpAdapter[RBT requests.IRequestBodyType] struct{}

func (ha *HttpAdapter[RBT]) Send(req requests.IRequest[RBT]) (responses.IResponse, error) {
	httpReq, ok := req.(*requests.HttpRequest[RBT])

	if ok {
		var bufferedBody *bytes.Buffer

		method  := httpReq.Method
		url     := httpReq.GetUri()
		headers := httpReq.Headers
		body    := httpReq.Body

		if reflect.TypeOf(body).String() == "string" {
			bufferedBody = bytes.NewBufferString(string(body))
		} else {
			bufferedBody = bytes.NewBuffer([]byte(body))
		}

		req, err := http.NewRequest(method, url, bufferedBody)

		if err != nil {
			panic(err.Error())
		}

		for name, value := range headers {
			req.Header.Add(name, value)
		}

		client := &http.Client{}
		res, err := client.Do(req)

		if err != nil {
			panic(err.Error())
		}

		defer res.Body.Close()

		return responses.NewResponse(responses.ParseHttp(res)), nil
	}

	return nil, errors.New("request of type HTTP required")
}
