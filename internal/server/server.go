package server

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"badger-api/pkg/server"
)

type BadgerServer struct {
	mux *http.ServeMux
}

func NewServer(ctx *server.ServerContext) BadgerServer {
	mux := http.NewServeMux()

	RegisterReflector(mux)

	RegisterActivityService(mux, ctx)
	RegisterDrillService(mux, ctx)

	return BadgerServer{
		mux,
	}
}

func (s *BadgerServer) Listen(addr string) {
	http.ListenAndServe(addr, h2c.NewHandler(s.mux, &http2.Server{}))
}
