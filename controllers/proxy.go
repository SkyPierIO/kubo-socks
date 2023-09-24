package controllers

import (
	"fmt"
	"strconv"

	socks5 "github.com/armon/go-socks5"
	"github.com/dProxSocks/kubo-socks/utils"
)

func StartProxy() {

	// Create a SOCKS5 server
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	config := utils.LoadConfiguration("./config.json")
	address := "127.0.0.1:" + strconv.Itoa(config.SocksPort)
	fmt.Println("SOCKS server is listening locally on " + strconv.Itoa(config.SocksPort))
	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", address); err != nil {
		panic(err)
	}
}
