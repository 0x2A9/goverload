**QUICK START**<br>
𝘨𝘙𝘗𝘊 𝘴𝘶𝘱𝘱𝘰𝘳𝘵 𝘪𝘴 𝘪𝘯 𝘥𝘦𝘷𝘦𝘭𝘰𝘱𝘮𝘦𝘯𝘵

```go
import (
	"github.com/lamia-mortis/goverload"
	"github.com/lamia-mortis/goverload/helpers/enums"
	"github.com/lamia-mortis/goverload/requests"
)
```

```go
/* HTTP Request */

params  := map[string]string{"key": "val"}
headers := map[string]string{"Content-Type": "application/json"}
body    := `{ "key1": [ [ 0, 0, 0, 0, 0, 0, 0, 0, 0 ] ], "key2": 3 }`

overloader      := goverload.NewOverloader[string]()
request         := goverload.NewRequest[string](enums.HTTP, "http-request-name-id")
httpRequest, ok := request.(*requests.HttpRequest[string])

if ok {
	httpRequest.
	    SetMethod("POST").
	    SetHost("host").
	    SetPort("8888").
	    SetPath("/api/some/path").
	    SetParams(params).
	    SetHeaders(headers).
	    SetBody(body)

	overloader.
	    AddRequest(httpRequest).
	    SetConfig(10, 10)

	overloader.Run()
}
```

```go
/* WS Request */

params  := map[string]string{"key": "val"}
headers := map[string]string{"Content-Type": "application/json"}
body    := `{ "key1": [ [ 0, 0, 0, 0, 0, 0, 0, 0, 0 ] ], "key2": 3 }`

overloader      := goverload.NewOverloader[string]()
request         := goverload.NewRequest[string](enums.WS, "ws-request-name-id")
wsRequest, ok   := request.(*requests.WsRequest[string])

if ok {
	wsRequest.
	    SetHost("host").
	    SetPort("8888").
	    SetPath("/api/some/path").
	    SetParams(params).
	    SetHeaders(headers).
	    SetBody(body)

	overloader.
	    AddRequest(wsRequest).
	    SetConfig(10, 10)

	overloader.Run()
}
```

```go
/* Adding multiple requests to the same run */

overloader   := goverload.NewOverloader[string]()

reqFirst     := goverload.NewRequest[string](enums.HTTP, "http-request-name-id")
httpReq, okF := req.(*requests.HttpRequest[string])

reqSecond    := goverload.NewRequest[string](enums.WS, "ws-request-name-id")
wsReq, okS   := reqSecond.(*requests.WsRequest[string])

if okF && okS {
	httpReq.
	    SetMethod("POST").
	    SetHost("host").
	    SetPort("8888")

	wsReq.
	    SetHost("host").
	    SetPort("88").
	    SetPath("/api/path")

	overloader.
	    AddRequest(httpReq).
	    SetConfig(10, 10)

	overloader.
	    AddRequest(wsReq).
	    SetConfig(5, 2)		

	overloader.Run()
}
```