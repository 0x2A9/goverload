package drivers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"

	"lamia-mortis/goverload/requests"
	"lamia-mortis/goverload/responses"
)

type WsAdapter[RBT requests.IRequestBodyType] struct{}

func (wa *WsAdapter[RBT]) Send(req requests.IRequest[RBT]) (responses.IResponse, error) {
	wsReq, ok := req.(*requests.WsRequest[RBT])

	if (ok) {
		conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), wsReq.GetOrigin())

		if err != nil {
			panic("Cannot connect: " + err.Error())
		}

		fmt.Println("Connected to server")

		var body []byte

		switch rawBody := any(wsReq.Body).(type) {
			case map[string]string:
			    body, err = json.Marshal(rawBody)

				if err != nil {
					panic("Request JSON parsind error: " + err.Error())
				}
			case string:
				body = []byte(rawBody)
		}

		err = wsutil.WriteClientMessage(conn, ws.OpText, body)

		if err != nil {
			panic("Cannot send: " + err.Error())
		}

		fmt.Println("Client message sent")

		res, _, err := wsutil.ReadServerData(conn)

		if err != nil {
			fmt.Printf("Cannot receive data: " + err.Error())
        }

		fmt.Println("Server message received")

		return wa.parseResponse(res), nil
	}

	return nil, errors.New("request of type WS required")
}

func (wa *WsAdapter[RBT]) parseResponse(res []byte) responses.IResponse {
	var content map[string]interface{}
	err := json.Unmarshal(res, &content)

	if err != nil {
		fmt.Printf("Response JSON parsind error: " + err.Error())
	}

	return &responses.Response{
		Body: content,
	}
}