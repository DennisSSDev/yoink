package main

import (
	"net/http"

	"github.com/dennisssdev/yoink/internal/greeterserver"
	"github.com/dennisssdev/yoink/rpc/greeter"
)

func main() {
	server := &greeterserver.Server{} // implements Haberdasher interface
	twirpHandler := greeter.NewGreeterServer(server, nil)

	http.ListenAndServe(":8034", twirpHandler)
}
