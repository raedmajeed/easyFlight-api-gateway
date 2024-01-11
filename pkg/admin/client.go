package admin

import (
	"log"

	pb "github.com/raedmajeed/api-gateway/pkg/admin/pb"
	"github.com/raedmajeed/api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ClientDial Client Dial method connect to the grpc admin client port
func ClientDial(cfg config.Configure) (pb.AdminAirlineClient, error) {
	grpcClient, err := grpc.Dial(":"+cfg.ADMINPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error dialing to grpc client: %s, ", cfg.ADMINPORT)
		return nil, err
	}
	log.Printf("succesfully connected to admin client at port: %v", cfg.ADMINPORT)
	return pb.NewAdminAirlineClient(grpcClient), nil
}
