package websock

import "github.com/gorilla/websocket"

//client represents a single websocket conn
type Client struct {
	Conn     *websocket.Conn
	PlayerID string
	GameID   string
}
