package api

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"github.com/gorilla/websocket"
	"github.com/newtoallofthis123/patients/db"
	"github.com/newtoallofthis123/patients/utils"
)

type ApiServer struct {
	store      *db.Store
	env        *utils.Env
	logger     *slog.Logger
	subs       map[*websocket.Conn]*genai.ChatSession
	sessionCtx map[*genai.ChatSession]context.Context
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
		store:      &store,
		env:        env,
		logger:     logger,
		subs:       make(map[*websocket.Conn]*genai.ChatSession),
		sessionCtx: make(map[*genai.ChatSession]context.Context),
	}, nil
}

func (s *ApiServer) Run() error {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

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
	r.POST("/transcribe", s.handleTranscribe)

	p := r.Group("/patients")
	p.POST("/create", s.handlePatientCreate)
	p.GET("/:id", s.handlePatientGet)
	p.PUT("/:id", s.handlePatientUpdate)

	err := r.Run(s.env.ListenAddr)
	return err
}
