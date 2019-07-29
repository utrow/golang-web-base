package server

import (
	"github.com/gin-gonic/gin"
	"github.com/utrow/golang-web-base/infrastructure/database/service"
	"github.com/utrow/golang-web-base/interface/router"
	"github.com/utrow/golang-web-base/registry"
	"log"
	"os"
)

func Launch() error {
	logger := log.New(os.Stdout, os.Getenv("GIT_REPOSITORY"), log.LstdFlags)

	db, err := service.CreateConnection()
	if err != nil {
		logger.Println(err)
	}

	r := registry.NewRegistry(db, logger)
	controller := r.NewController()
	server := gin.Default()

	server.Use(Cors)

	router.NewRouter(server, controller)

	// listen and serve on 0.0.0.0:8080
	return server.Run()
}

func Cors(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.Next()
}
