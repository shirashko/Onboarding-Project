package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirrashko/BuildingAServer-step2/config"
)

type Server struct {
	engine *gin.Engine
	conf   config.Config
}

func NewServer(conf config.Config, routerFactory func(config.Config) (Handlers, error)) (Server, error) {
	engine := gin.Default() // The default gin use Logger & Recovery middlewares
	server := Server{engine: engine, conf: conf}
	handlers, err := routerFactory(conf)
	if err != nil {
		return server, err
	}
	server.SetUp(handlers)
	return server, nil
}

func (server *Server) SetUp(handlers Handlers) {
	for _, h := range handlers.handlers { // looping over a list of handlers with idx and handler=value. Each handler
		// implements its own setup routes function according to its API
		h.SetUpRoutes(server.engine)
	}
}

func (server *Server) ListenAndServe() error {
	address := fmt.Sprintf("%s:%d", server.conf.ServerInfo.Host, server.conf.ServerInfo.Port)
	return server.engine.Run(address) // tell the engine to listen and serve localhost:8080
}
