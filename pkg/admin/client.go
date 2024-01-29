package admin

import (
	pb "github.com/raedmajeed/api-gateway/pkg/admin/pb"
	"github.com/raedmajeed/api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ClientDial Client Dial method connect to the grpc admin client port
func ClientDial(cfg config.Configure) (pb.AdminAirlineClient, error) {
	grpcClient, err := grpc.Dial(cfg.ADMINPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewAdminAirlineClient(grpcClient), nil
}
