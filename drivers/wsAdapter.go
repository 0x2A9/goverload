package drivers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"

	"lamia-mortis/goverload/requests"
	"lamia-mortis/goverload/responses"
)

type WsAdapter[RBT requests.IRequestBodyType] struct{}

func (wa *WsAdapter[RBT]) Send(req requests.IRequest[RBT]) (responses.IResponse, error) {
	wsReq, ok := req.(*requests.WsRequest[RBT])

	if ok {
		conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), wsReq.GetUri())

		if err != nil {
			panic("Cannot connect: " + err.Error())
		}

		fmt.Println("Connected to server")

		var body []byte

		// using `any` for having ability to check the generic parameter type
		switch rawBody := any(wsReq.Body).(type) {
		case map[string]string:
			body, err = json.Marshal(rawBody)

			if err != nil {
				panic("Request JSON parsing error: " + err.Error())
			}
		case string:
			body = []byte(rawBody)
		}

		err = wsutil.WriteClientMessage(conn, ws.OpText, body)

		if err != nil {
			panic("Cannot send: " + err.Error())
		}

		fmt.Println("Client message sent")

		wsRes, opCode, err := wsutil.ReadServerData(conn)

		if err != nil {
			fmt.Printf("Cannot receive data: " + err.Error())
		}

		fmt.Println("Server message received")

		err = conn.Close()

		if err != nil {
			fmt.Println("Cannot close the connection: " + err.Error())
			os.Exit(1)
		}

		fmt.Println("Disconnected from server")

		return responses.NewResponse(responses.ParseWs(opCode, wsRes)), nil
	}

	return nil, errors.New("request of type WS required")
}

