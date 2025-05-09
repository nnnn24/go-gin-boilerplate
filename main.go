package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nnnn24/go-gin-boilerplate/pkg/conf"
	"github.com/nnnn24/go-gin-boilerplate/pkg/logging"
	"github.com/nnnn24/go-gin-boilerplate/routers"
)

func init() {
	logging.Setup()
}

func main() {
	cfg := conf.Get()

	gin.SetMode(cfg.GinRunMode)

	routersInit := routers.Init()
	endPoint := fmt.Sprintf(":%s", cfg.ServerPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    60 * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	err := server.ListenAndServe()

	if err != nil {
		log.Printf("Server err: %v", err)
	}

}
