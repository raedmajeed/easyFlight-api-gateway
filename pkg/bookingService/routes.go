package bookingservice

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raedmajeed/api-gateway/middleware"
	"github.com/raedmajeed/api-gateway/pkg/bookingService/handlers"
	pb "github.com/raedmajeed/api-gateway/pkg/bookingService/pb"
	"github.com/raedmajeed/api-gateway/pkg/config"
)

type BookingServer struct {
	cfg *config.Configure
	pb  pb.BookingClient
}

func NewBookingRoutes(ctx *gin.Engine, cfg config.Configure) {
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error Not connected with gRPC server, %v", err.Error())
	}
	bookingServer := BookingServer{
		cfg: &cfg,
		pb:  client,
	}

	apiVersion := ctx.Group("/api/v1")
	{
		flights := apiVersion.Group("/flights")
		{
			flights.GET("/search", bookingServer.SearchFlight)
		}

	}
}

func (bs *BookingServer) UserAuthenticate(ctx *gin.Context) {
	email, err := middleware.ValidateToken(ctx, *bs.cfg, "user")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":  err.Error(),
			"status": http.StatusUnauthorized,
		})
		return
	}
	ctx.Set("registered_email", email)
	ctx.Next()
}

func (bs *BookingServer) SearchFlight(ctx *gin.Context) {
	handlers.SearchFlight(ctx, bs.pb)
}
