# kubo-socks-plugin
A Kubo plugin allowing to use SOCKS proxy over the IPFS network

## Build

You can build the plugin using 

```
go build . -o kubo-socks
```

And then run it with 

```
# Be sure to have a kubo node already running
./kubo-socks
```

## SOCKS5 Proxy

The plugin is running a local SOCKS5 proxy on port `tcp/1080`.
It is not supposed to be used as it, but buy a remote client coming from another `kubo-socks` node.

## API

The plugin serves an HTTP API that can be requested from the frontend

- GET `/ping`  ↔ Just ping the backend
- GET `/streams`  ↔ List all active libp2p streams on the Kubo node
- GET `/listeners`  ↔ List all listeners on the Kubo node
- GET `/peers`  ↔ Show directly connected peers of Kubo node
- GET `/forward/<nodeID>`  ↔ Open a port locally listening to SOCKS5 clients and forwarding connections to the node <nodeID>
- GET `/ping/<nodeID>`  ↔ Send echo request packets to IPFS host <nodeID>
- GET `/streams/close`  ↔ Close **all** libp2p streams on the Kubo node
- GET `/id`  ↔ Return the local nodeID 

## We're hacking on IPFS <3

[![](https://cdn.rawgit.com/jbenet/contribute-ipfs-gif/master/img/contribute.gif)](https://github.com/ipfs/community/blob/master/CONTRIBUTING.md)