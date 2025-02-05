package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func (s *ApiServer) handleTalk(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error upgrading websocket: %v", err)
		return
	}

	s.subs[conn] = make([]int, 0)

	for {
		_, prompt, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		if string(prompt) == "bye" {
			delete(s.subs, conn)
			c.JSON(200, gin.H{"success": "Response generated"})
			return
		}

		userCtx := s.subs[conn]
		res, err := s.talk(string(prompt), userCtx)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		s.subs[conn] = res.Context

		conn.WriteMessage(websocket.TextMessage, []byte(res.Response))
	}
}
