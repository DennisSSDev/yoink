package epicstorecommunicatorserver

import (
	"context"

	pb "github.com/dennisssdev/yoink/rpc/epicstorecommunicator"
)

/*
IsAlive is a ping function to be used to check
if the server can make a simple response.
It is mainly used to detect if the server has shut down completely
*/
func (s *Server) IsAlive(ctx context.Context, empty *pb.EmptyReq) (resp *pb.IsAliveResp, err error) {
	return &pb.IsAliveResp{
		Status: "The service is up",
	}, nil
}
