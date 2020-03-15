package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ibinarytree/koala/server"

	"github.com/ibinarytree/hello/router"

	"github.com/ibinarytree/hello/proto_gen/hello"
)

var routerServer = &router.RouterServer{}
var (
	BUILD_TIME string
	GO_VERSION string
	GIT_COMMIT string
)

func main() {

	if len(os.Args) >= 2 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Printf("build time:%s\n", BUILD_TIME)
		fmt.Printf("go version:%s\n", GO_VERSION)
		fmt.Printf("git commit:%s\n", GIT_COMMIT)
		return
	}

	err := server.Init("hello")
	if err != nil {
		log.Fatal("init service failed, err:%v", err)
		return
	}

	hello.RegisterHelloServiceServer(server.GRPCServer(), routerServer)
	server.Run()
}
