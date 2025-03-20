package websock

import (
	"encoding/json"
	"log"
	"net/http"
	"tic-tac-toe/game"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Upgrader handles websocket Upgrades
var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, //allow all origin
}

//Message defines incomming Websocket Messages

type Message struct {
	Action string `json:"action"`
	GameID string `json:"game_id"`
	Player string `json:"player"`
	Row    int    `json:"row"`
	Col    int    `josn:"col"`
}

type Handler struct {
	Hub   *Hub
	Games *game.Manager
}

//New Handler initializes the Websocket Handler

func NewHandler(hub *Hub, game *game.Manager) *Handler {
	return &Handler{
		Hub:   hub,
		Games: game,
	}
}

// Websocket Handler handles new Websocket connection
func (h *Handler) Websockethandler(c *gin.Context) {
	conn, err := Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Websocket upgrader error", err)
		return
	}
	defer conn.Close()

	client := &Client{Conn: conn}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error :", err)
			h.Hub.UnregisterAClient(client.PlayerID)
			break
		}

		var message Message
		err = json.Unmarshal(msg, &message)
		if err != nil {
			log.Println("Invalid message format", err)
			continue
		}
		h.HandleMessage(client, message)
	}
}

//Handles message actions(join,move, etc)

func (h *Handler) HandleMessage(client *Client, message Message) {
	switch message.Action {
	case "join":
		h.HandleJoin(client, message)
	case "move":
		h.HandleMove(client, message)
	default:
		client.Conn.WriteJSON(map[string]string{"error": "Invalid action"})
	}
}

// Handle player joining a Game
func (h *Handler) HandleJoin(client *Client, msg Message) {
	gameId := msg.GameID
	playerID := msg.Player

	game, exist := h.Games.Games[gameId]
	if !exist {
		_, err := h.Games.CreateGame(gameId, playerID)
		if err != nil {
			client.Conn.WriteJSON(map[string]string{"error:": err.Error()})
			return
		}
	} else {
		err := game.AddPlayer(playerID)
		if err != nil {
			client.Conn.WriteJSON(map[string]string{"error:": err.Error()})
			return
		}
	}
	client.GameID = gameId
	client.PlayerID = playerID
	h.Hub.RegisterNewClient(client)

	client.Conn.WriteJSON(map[string]string{"message": "Successfully joined the game"})
}

//handle moves

func (h *Handler) HandleMove(client *Client, msg Message) {
	games, exist := h.Games.Games[msg.GameID]
	if !exist {
		client.Conn.WriteJSON(map[string]string{"error:": "GAme does not exist"})
		return
	}
	err := games.MakeAMove(msg.Row, msg.Col, games.NextTurn)
	if err != nil {
		client.Conn.WriteJSON(map[string]string{"error:": err.Error()})
		return
	}

	//check Win Condition

	winner := games.Board.CheckWinner()
	if winner != "" {
		h.Hub.Brodcast(msg.GameID, map[string]string{
			"status": "win",
			"winner": winner,
		})
		h.Games.EndGame(msg.GameID)
		return
	}

	//check draw condition
	if games.Board.CheckDraw() {
		h.Hub.Brodcast(msg.GameID, map[string]string{
			"status": "draw",
		})
		h.Games.EndGame(msg.GameID)
		return
	}

	//Brodcast updated Board status to both players

	h.Hub.Brodcast(msg.GameID, map[string]interface{}{
		"status": "update",
		"board":  games.Board.Grid,
		"next":   games.NextTurn,
	})
}
