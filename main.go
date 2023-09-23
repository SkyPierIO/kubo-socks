package main

import (
	"strconv"

	"github.com/dProxSocks/kubo-socks-plugin/controllers"
	"github.com/dProxSocks/kubo-socks-plugin/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	config := utils.LoadConfiguration("./config.json")
	router := gin.Default()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	api := router.Group("/api/v0")
	api.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	api.GET("/stream/ls", controllers.ListStreams)
	api.GET("/listener/ls", controllers.ListListeners)
	api.GET("/peers", controllers.ShowPeers)
	api.GET("/forward", controllers.Forward)

	// Run with HTTP
	router.Run("0.0.0.0:" + strconv.FormatUint(uint64(config.Port), 10))
}
