package api

import (
	"fmt"
	"log"
	"schule/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	engine *gin.Engine
	config *config.Config
	db     *gorm.DB
}

func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	r := gin.Default()

	s := &Server{
		engine: r,
		config: cfg,
		db:     db,
	}

	s.registerRoutes()
	return s
}

func (s *Server) registerRoutes() {
	RegisterRoutes(s.engine, s.db)
}

func (s *Server) Run() {
	addr := fmt.Sprintf(":%s", s.config.Port)
	log.Printf("Server running on %s", addr)
	s.engine.Run(addr)
}
