package main

import (
	"os"
	"github.com/pborman/uuid"
	"strings"
	"fmt"
)

func main() {
	mode := "client"

	if len(os.Args) > 1 && os.Args[1] == "--server" {
		mode = "server"
	} else if len(os.Args) > 1 && os.Args[1] == "--help" {
		mode = "help"
	}

	switch mode {
	case "server":
		ip, port, sslCrt, sslKey := ParseServerFlags(os.Args[2:])
		RunServer(ip, port, sslCrt, sslKey, false)
	case "client":
		url, userKey, rootPath := ParseClientFlags(os.Args[1:])
		id := strings.Replace(uuid.New(), "-", "", -1)
		RunClient(url, id, userKey, rootPath)
	case "help":
		fmt.Println(`zedrem runs in one of two possible modes: client or server:

Usage: zedrem [-u url] [-key userKey] <dir>
       Launches a Zed client and attaches to a Zed server exposing
       directory <dir> (or current directory if omitted). Default URL is
       wss://remote.zedapp.org:443
       If a -key flag is passed that matches the userKey set in your Zed
       configuration, a window will open automatically.

Usage: zedrem --server [-h ip] [-p port] [--sslcrt file.crt] [--sslkey file.key]
       Launches a Zed server, binding to IP <ip> on port <port>.
       If --sslcrt and --sslkey are provided, will run in TLS mode for more security.
`)
	}
}
