package handlers

import (
	"fmt"
	"log"
)

type WsPayload struct {
	Action  string              `json:"action"`
	Message string              `json:"message"`
	Conn    WebSocketConnection `json:"-"`
}

type WsMessage struct {
	Payload *WsPayload           // the JSON payload
	Conn    *WebSocketConnection // pointer to the live connection
}

// ListenForWs : listen for user question
func ListenForWs(conn *WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Errror", fmt.Sprintf("%v", r))
		}
	}()

	var payload WsPayload

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			// do nothing
			log.Println("ws read err", err)
			break
		} else {
			wsMessage := WsMessage{
				Payload: &payload,
				Conn:    conn,
			}
			fmt.Println("Sending payload to channel", wsMessage)
			wsChan <- wsMessage // send payload to channel
		}
	}
}

func ListenToWsChannel() {
	//var response WsJsonResponse

	for {
		e := <-wsChan // read payload from channel
		fmt.Println("listning fo webhook event")
		switch e.Payload.Action {

		case "question":
			//get question
			fmt.Println("message from gondor", e.Payload.Action, e.Payload.Message)

		}
	}
}
