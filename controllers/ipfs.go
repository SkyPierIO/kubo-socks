package controllers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/SkyPierIO/kubo-socks/utils"
	"github.com/gin-gonic/gin"
	shell "github.com/ipfs/go-ipfs-api"
)

// ----------------------------------------------------------------------------
// 		   _____ __                  __
// 		  / ___// /________  _______/ /_
// 		  \__ \/ __/ ___/ / / / ___/ __/
// 		 ___/ / /_/ /  / /_/ / /__/ /_
// 		/____/\__/_/   \__,_/\___/\__/
//
// ----------------------------------------------------------------------------

type P2PListenerList struct {
	Listeners []*P2PListener
}

type P2PListener struct {
	Protocol      string
	ListenAddress string
	TargetAddress string
}

type P2PStream struct {
	Protocol string
	Address  string
}

type Redirection struct {
	Protocol string
	Target   string
}

type PingResult struct {
	Success bool
	Text    string
	Time    int
}

type ConfigResult struct {
	Key   string
	Value bool
}

type P2PStreamsList struct {
	Streams []*struct {
		HandlerID     string
		Protocol      string
		LocalPeer     string
		LocalAddress  string
		RemotePeer    string
		RemoteAddress string
	}
}

type Node struct {
	ID string
}

type Peer struct {
	Addr string
	Peer string
}

type PeerList struct {
	Peers []*Peer
}

type ForwardResponse struct {
	Success          bool
	Status           string
	ListeningAddress string
	ListeningPort    string
}

// -------------------------------------------------------------------
// 		    __  ___     __  __              __
// 		   /  |/  /__  / /_/ /_  ____  ____/ /____
// 		  / /|_/ / _ \/ __/ __ \/ __ \/ __  / ___/
// 		 / /  / /  __/ /_/ / / / /_/ / /_/ (__  )
// 		/_/  /_/\___/\__/_/ /_/\____/\__,_/____/
//
// -------------------------------------------------------------------

func ListListeners(c *gin.Context) {
	var response *P2PListenerList
	s := shell.NewLocalShell()
	reqBuilder := s.Request("p2p/ls")
	err := reqBuilder.Exec(context.Background(), &response)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, response.Listeners)
}

func ListStreams(c *gin.Context) {
	var response *P2PStreamsList
	s := shell.NewLocalShell()
	reqBuilder := s.Request("p2p/stream/ls")
	err := reqBuilder.Exec(context.Background(), &response)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, response.Streams)
}

func EnableLibp2pStreaming() {
	var response *ConfigResult
	s := shell.NewLocalShell()
	reqBuilder := s.Request("config")
	reqBuilder.Arguments("Experimental.Libp2pStreamMounting", "true")
	reqBuilder.Option("bool", true)
	err := reqBuilder.Exec(context.Background(), &response)
	if err != nil {
		log.Fatal(err)
	}
}

// Send echo request packets to IPFS hosts.
func Ping(c *gin.Context) {
	var response *PingResult
	nodeID := c.Param("nodeID")
	s := shell.NewLocalShell()
	reqBuilder := s.Request("ping")
	reqBuilder.Arguments(nodeID, "1") // args: nodeId, ping count
	err := reqBuilder.Exec(context.Background(), &response)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, response)
}

func Forward(c *gin.Context) {
	var response string
	nodeID := c.Param("nodeID")
	port := utils.GetFirstAvailablePort()
	customProtocol := "/x/7proxies/1.0"
	listenAddr := "/ip4/127.0.0.1/tcp/" + port
	fmt.Println(listenAddr)
	s := shell.NewLocalShell()
	reqBuilder := s.Request("p2p/forward")
	reqBuilder.Arguments(customProtocol, listenAddr, "/p2p/"+nodeID)
	err := reqBuilder.Exec(context.Background(), &response)
	if err != nil {
		if err == io.EOF {
			resp := ForwardResponse{true, "Forward enabled to node " + nodeID, "127.0.0.1", port}
			c.IndentedJSON(http.StatusOK, resp)
		} else {
			fmt.Println(err)
			resp := ForwardResponse{false, "Cannot enable Forward for this connection: " + err.Error(), "0", "0"}
			c.IndentedJSON(http.StatusInternalServerError, resp)
		}
	}
}

func Listen(protocol string, target string) {
	var response string
	redirection := Redirection{
		Protocol: protocol,
		Target:   target,
	}
	s := shell.NewLocalShell()
	reqBuilder := s.Request("p2p/listen")
	reqBuilder.Arguments(redirection.Protocol, redirection.Target)
	err := reqBuilder.Exec(context.Background(), &response)
	if err != nil {
		if err == io.EOF {
			fmt.Println("Listening for new proxy connection")
		} else {
			fmt.Println(err)
		}
	}
}

func CloseAllSteams(c *gin.Context) {
	var response int
	s := shell.NewLocalShell()
	reqBuilder := s.Request("p2p/close")
	reqBuilder.Option("all", true)
	err := reqBuilder.Exec(context.Background(), &response)
	if err != nil {
		if err == io.EOF {
			c.String(http.StatusOK, "OK")
		} else {
			log.Fatal(err)
		}
	}

}

func ShowPeers(c *gin.Context) {
	var response *PeerList
	s := shell.NewLocalShell()
	reqBuilder := s.Request("swarm/peers")
	err := reqBuilder.Exec(context.Background(), &response)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, response.Peers)
}

func GetID(c *gin.Context) {
	var response *Node
	s := shell.NewLocalShell()
	reqBuilder := s.Request("id")
	err := reqBuilder.Exec(context.Background(), &response)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, response.ID)
}
