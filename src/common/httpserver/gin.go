package httpserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/mcuadros/go-gin-prometheus"
)

type HTTPServer struct {
	Router *gin.Engine
	port   int
	srv    *http.Server
}

type Config struct {
	Port int
}

func New() *HTTPServer {
	conf := GetConfig()

	router := gin.New()

	p := ginprometheus.NewPrometheus("gin")
	p.SetListenAddressWithRouter(fmt.Sprintf("0.0.0.0:%d", 2100), gin.New())
	p.Use(router)

	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	return &HTTPServer{
		port:   conf.Port,
		Router: router,
	}
}

func GetConfig() Config {
	strport := os.Getenv("PORT")

	port, err := strconv.Atoi(strport)
	if err != nil {
		log.Fatalf("Incorrect port %s", strport)
	}

	return Config{
		Port: port,
	}
}

func (h *HTTPServer) ListenAndServe() {
	h.srv = &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", h.port),
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 60,
		Handler:      h.Router,
	}

	if err := h.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
func (h *HTTPServer) Shutdown() {
	log.Println("Shutdown HTTP Server...")

	if err := h.srv.Shutdown(context.Background()); err != nil {
		log.Printf("server Shutdown: %v", err)
	}
}
