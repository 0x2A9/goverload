package drivers

import (
	"encoding/json"
	"errors"
	"net/http"

	"lamia-mortis/goverload/requests"
	"lamia-mortis/goverload/responses"
)

type HttpAdapter[RBT requests.IRequestBodyType] struct {}

func (ha *HttpAdapter[RBT]) Send(req requests.IRequest[RBT]) (responses.IResponse, error) {
	httpReq, ok := req.(*requests.HttpRequest[RBT])

	if (ok) {
		method  := httpReq.Method
		url     := httpReq.GetOrigin() + httpReq.Path
		headers := httpReq.Headers
		// body    := httpReq.Body

		req, err := http.NewRequest(method, url, nil)

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

		return ha.parseResponse(res), nil
	}

	return nil, errors.New("Request of type HTTP required")
}

func (ha *HttpAdapter[RBT]) parseResponse(res *http.Response) responses.IResponse {
    var content map[string]interface{}
    json.NewDecoder(res.Body).Decode(&content)

	return &responses.Response{
		Body: content,
	}
}