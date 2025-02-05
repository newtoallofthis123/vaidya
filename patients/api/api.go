package api

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/newtoallofthis123/patients/db"
	"github.com/newtoallofthis123/patients/utils"
)

type ApiServer struct {
	store  *db.Store
	env    *utils.Env
	logger *slog.Logger
	subs   map[*websocket.Conn][]int
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewApiServer(env *utils.Env) (*ApiServer, error) {
	store, err := db.NewStore(env.ConnString)
	if err != nil {
		return nil, err
	}

	err = store.InitTables()
	if err != nil {
		return nil, err
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	return &ApiServer{
		store:  &store,
		env:    env,
		logger: logger,
		subs:   make(map[*websocket.Conn][]int),
	}, nil
}

func (s *ApiServer) Run() error {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/talk", s.handleTalk)

	err := r.Run(s.env.ListenAddr)
	return err
}
