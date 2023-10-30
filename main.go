package main

import (
	"fmt"
	"strconv"

	"github.com/dProxSocks/kubo-socks/controllers"
	"github.com/dProxSocks/kubo-socks/utils"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {

	config := utils.LoadConfiguration("./config.json")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())

	fmt.Println("───────────────────────────────────────────────────")
	fmt.Println("🌎😎 ~~ YOU ARE RUNNING KUBO SOCKS PLUGIN ~~  😎🌎")
	fmt.Println("           ~~ let's browse the GALAXY ~~      ")
	fmt.Println("───────────────────────────────────────────────────")

	go func() {
		// http.ListenAndServe("localhost:8081", serverMuxA)
		controllers.StartProxy()
	}()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	router.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})

	api := router.Group("/api/v0")
	api.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})
	api.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	api.GET("/streams", controllers.ListStreams)
	api.GET("/listeners", controllers.ListListeners)
	api.GET("/peers", controllers.ShowPeers)
	api.GET("/forward/:nodeID", controllers.Forward)
	api.GET("/ping/:nodeID", controllers.Ping)
	api.GET("/streams/close", controllers.CloseAllSteams)
	api.GET("/id", controllers.GetID)

	// Enable the Listener by default on the proxy port
	protocol := "/x/7proxies/1.0"
	target := "/ip4/127.0.0.1/tcp/" + strconv.Itoa(config.SocksPort)
	controllers.EnableLibp2pStreaming()
	controllers.Listen(protocol, target)

	fmt.Println("───────────────────────────────────────────────────")

	// Run with HTTP
	router.Run("0.0.0.0:" + strconv.FormatUint(uint64(config.Port), 10))
}
