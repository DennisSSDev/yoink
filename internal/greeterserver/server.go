package greeterserver

import (
	"context"

	"github.com/twitchtv/twirp"

	pb "github.com/dennisssdev/yoink/rpc/greeter"
)

// Server implements the Greeter service
type Server struct{}

// MakeGreeting creates a simple greeting response
func (s *Server) MakeGreeting(ctx context.Context, name *pb.NameInput) (hat *pb.GreeterResponse, err error) {
	if name.GetName() == "" {
		return nil, twirp.InvalidArgumentError("Name", "must be provided and none-blank")
	}
	return &pb.GreeterResponse{
		Text: "HI " + name.GetName(),
	}, nil
}
