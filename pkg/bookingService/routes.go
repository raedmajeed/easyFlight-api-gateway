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
			search := flights.Group("/search")
			{
				search.GET("", bookingServer.SearchFlight)
				search.GET("/select", bookingServer.SelectFlight)
				search.POST("/select/pax", bookingServer.UserAuthenticate, bookingServer.AddTravellers)
			}
		}
		user := apiVersion.Group("/user")
		{
			user.POST("/login", bookingServer.UserLogin)
			user.POST("/register", bookingServer.UserRegister)
			user.POST("/register/verify", bookingServer.VerifyRegistration2)
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

func (bs *BookingServer) SelectFlight(ctx *gin.Context) {
	handlers.SelectFlight(ctx, bs.pb)
}

func (bs *BookingServer) AddTravellers(ctx *gin.Context) {
	handlers.AddTravellers(ctx, bs.pb)
}

func (bs *BookingServer) UserLogin(ctx *gin.Context) {
	handlers.UserLogin(ctx, bs.pb)
}
func (bs *BookingServer) UserRegister(ctx *gin.Context) {
	handlers.UserRegister(ctx, bs.pb)
}
func (bs *BookingServer) VerifyRegistration2(ctx *gin.Context) {
	handlers.VerifyRegistration2(ctx, bs.pb)
}
