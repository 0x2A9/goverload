package responses

import (
	"io"
	"net/http"

	"github.com/gobwas/ws"
)

func ParseHttp(res *http.Response) (map[string]any, []byte) {
	headers := make(map[string]any)

	for name := range res.Header {
		headers[name] = res.Header.Get(name)
	}

	content, error := io.ReadAll(res.Body)

	if error != nil {
	   panic("Error during parsing response: " + error.Error())
	}

	return headers, content
}

func ParseWs(opCode ws.OpCode, body []byte) (map[string]any, []byte) {
	headers := make(map[string]any)

	headers["opCode"] = opCode

	return headers, body
}
