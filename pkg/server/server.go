package server

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/polyrain/go-keno/pkg"
	slogecho "github.com/samber/slog-echo"
)

// Only one server should ever exist
var singleton *Server

// Our server definition
type Server struct {
	cfg *pkg.ServerConfig
}

func NewServer(cfg *pkg.ServerConfig) *Server {
	singleton = &Server{cfg: cfg}
	return singleton
}

func GetServer() *Server {
	return singleton
}

func (s *Server) Start() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())

	// Use slog as our logging soln for the server
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	e.Use(slogecho.New(logger))

	// Routes here

	// Start up the server and log if theres any issues
	e.Logger.Fatal(e.Start(s.cfg.GetPort()))
}
