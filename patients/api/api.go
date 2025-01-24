package api

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/newtoallofthis123/patients/db"
	"github.com/newtoallofthis123/patients/utils"
)

type ApiServer struct {
	store  *db.Store
	env    *utils.Env
	logger *slog.Logger
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
	}, nil
}

func (s *ApiServer) Run() error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run(s.env.ListenAddr)
	return err
}
