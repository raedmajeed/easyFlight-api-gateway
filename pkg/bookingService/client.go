package bookingservice

import (
	"log"

	pb "github.com/raedmajeed/api-gateway/pkg/bookingService/pb"
	"github.com/raedmajeed/api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Configure) (pb.BookingClient, error) {
	grpc, err := grpc.Dial(":"+cfg.BSERVICEPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc client: %s, ", cfg.BSERVICEPORT)
		return nil, err
	}
	log.Printf("succesfully Connected to Booking Client at port: %v", cfg.ADMINPORT)
	return pb.NewBookingClient(grpc), nil
}