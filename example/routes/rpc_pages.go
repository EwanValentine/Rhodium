package routes

import (
	"rhodium"
	"rhodium/example/rpc"
)

type PostRPCRoutes struct {
	rpcController *rpc.PostRPC
}

// NewPostRPCRoutes -
func NewPostRPCRoutes(rpcController *rpc.PostRPC) *PostRPCRoutes {
	return &PostRPCRoutes{rpcController}
}

// CreatePost -
func (r *PostRPCRoutes) CreatePost() rhodium.RPCRoute {
	return rhodium.RPCRoute{
		Name:    "create.post",
		Handler: r.rpcController.CreatePost,
	}
}

// Routes -
func (r *PostRPCRoutes) Routes() []rhodium.RPCRoute {
	return []rhodium.RPCRoute{
		r.CreatePost(),
	}
}
