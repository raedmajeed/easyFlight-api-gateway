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
			user.POST("/bookings/confirm", bookingServer.UserAuthenticate, bookingServer.ConfirmBooking)
			user.POST("/booking/confirm/online/payment", bookingServer.UserAuthenticate, bookingServer.OnlinePayment)
			user.POST("/booking/confirm/online/payment/success", bookingServer.UserAuthenticate, bookingServer.PaymentSuccess)
			user.POST("/booking/confirm/online/payment/success/render", bookingServer.UserAuthenticate, bookingServer.PaymentSuccessPage)
			confirmedUser := user.Group("/confirmed")
			{
				confirmedUser.POST("/login", bookingServer.PNRLogin)
				confirmedUser.POST("/selectSeats", bookingServer.UserPNRAuthenticate, bookingServer.SelectSeat)
			}
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

func (bs *BookingServer) UserPNRAuthenticate(ctx *gin.Context) {
	pnr, err := middleware.ValidatePNRToken(ctx, *bs.cfg, "PNR-USER")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":  err.Error(),
			"status": http.StatusUnauthorized,
		})
		return
	}
	ctx.Set("user_pnr", pnr)
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

func (bs *BookingServer) SelectSeat(ctx *gin.Context) {
	handlers.SelectSeat(ctx, bs.pb)
}

func (bs *BookingServer) PNRLogin(ctx *gin.Context) {
	handlers.PNRLogin(ctx, bs.pb)
}

func (bs *BookingServer) ConfirmBooking(ctx *gin.Context) {
	handlers.ConfirmBooking(ctx, bs.pb)
}

func (bs *BookingServer) OnlinePayment(ctx *gin.Context) {
	handlers.OnlinePayment(ctx, bs.pb)
}

func (bs *BookingServer) PaymentSuccess(ctx *gin.Context) {
	handlers.PaymentSuccess(ctx, bs.pb)
}
func (bs *BookingServer) PaymentSuccessPage(ctx *gin.Context) {
	handlers.PaymentSuccessPage(ctx, bs.pb)
}
