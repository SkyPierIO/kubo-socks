package controllers

import (
	"context"
	"io"
	"log"
	"net/http"

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
	// reqBuilder.Arguments("arg1", "arg2")
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
	// reqBuilder.Arguments("arg1", "arg2")
	err := reqBuilder.Exec(context.Background(), &response)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, response.Streams)
}

func Forward(c *gin.Context) {
	var response string
	s := shell.NewLocalShell()
	reqBuilder := s.Request("p2p/forward")
	reqBuilder.Arguments("/x/7proxies/1.0", "/ip4/127.0.0.1/tcp/3333", "/p2p/12D3KooWNngZko3MYLgd73W6MspMQdYWkcXP2dqBHWZR8nheLnpG")
	err := reqBuilder.Exec(context.Background(), &response)
	if err != nil {
		if err == io.EOF {
			c.String(http.StatusOK, "OK")
		} else {
			c.String(http.StatusInternalServerError, "Cannot enable Forward for this connection")
		}
	}

}

func CloseAllSteams(c *gin.Context) {
	var response string
	s := shell.NewLocalShell()
	reqBuilder9050 := s.Request("p2p/close")
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
