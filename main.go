package main

import (
	"log"
	"tic-tac-toe/game"
	routes "tic-tac-toe/router"
	websock "tic-tac-toe/websocket"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//initialize Websocket components
	hub := websock.NewHub()
	gamemanager := game.NewManager()
	wshandler := websock.NewHandler(hub, gamemanager)

	//setuproutes
	routes.Steproutes(r, wshandler)

	//Run server

	port := ":8080"
	log.Println("Server is running on port ", port)
	log.Fatal(r.Run(port))
}
