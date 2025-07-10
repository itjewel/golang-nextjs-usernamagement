package server

import (
    "github.com/gin-gonic/gin"
    "user-service/internal/handler"
)

type Server struct {
    engine *gin.Engine
}

func NewServer() *Server {
    return &Server{engine: gin.Default()}
}

func (s *Server) RegisterRoutes(h *handler.UserHandler) {
    s.engine.GET("/users", h.GetUsers)
}

func (s *Server) Run(addr string) error {
    return s.engine.Run(addr)
}
