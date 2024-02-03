package pkg

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/raedmajeed/api-gateway/docs"
	"github.com/raedmajeed/api-gateway/middleware"
	c "github.com/raedmajeed/api-gateway/pkg/bookingService"
	"github.com/raedmajeed/api-gateway/pkg/bookingService/handlers"
	pb "github.com/raedmajeed/api-gateway/pkg/bookingService/pb"
	"github.com/raedmajeed/api-gateway/pkg/config"
)

type BookingServer struct {
	cfg *config.Configure
	pb  pb.BookingClient
}

func NewBookingRoutes(ctx *gin.Engine, cfg config.Configure) {
	client, err := c.ClientDial(cfg)
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
			user.GET("/booking/confirm/online/payment", bookingServer.OnlinePayment)
			user.GET("/booking/confirm/online/payment/success", bookingServer.PaymentSuccess)
			user.GET("/booking/confirm/online/payment/success/render", bookingServer.PaymentSuccessPage)
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

// @Summary Search for a flight
// @Description Search for a flight
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/booking/search-flight [get]
// @Tags FlightSearch
func (bs *BookingServer) SearchFlight(ctx *gin.Context) {
	handlers.SearchFlight(ctx, bs.pb)
}

// @Summary Select a flight
// @Description Select a flight
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/booking/select-flight [post]
// @Tags FlightSearch
func (bs *BookingServer) SelectFlight(ctx *gin.Context) {
	handlers.SelectFlight(ctx, bs.pb)
}

// @Summary Add travellers to a flight
// @Description Add travellers to a flight
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/booking/add-travellers [post]
// @Tags FlightSearch
func (bs *BookingServer) AddTravellers(ctx *gin.Context) {
	handlers.AddTravellers(ctx, bs.pb)
}

// @Summary User login
// @Description User login
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/booking/user-login [post]
// @Tags User
func (bs *BookingServer) UserLogin(ctx *gin.Context) {
	handlers.UserLogin(ctx, bs.pb)
}

// @Summary User registration
// @Description User registration
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/booking/user-register [post]
// @Tags User
func (bs *BookingServer) UserRegister(ctx *gin.Context) {
	handlers.UserRegister(ctx, bs.pb)
}

// @Summary Verify user registration
// @Description Verify user registration
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/booking/verify-registration [post]
// @Tags User
func (bs *BookingServer) VerifyRegistration2(ctx *gin.Context) {
	handlers.VerifyRegistration2(ctx, bs.pb)
}

// @Summary Select a seat
// @Description Select a seat
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/booking/select-seat [post]
// @Tags ConfirmedUser
func (bs *BookingServer) SelectSeat(ctx *gin.Context) {
	handlers.SelectSeat(ctx, bs.pb)
}

// @Summary PNR login
// @Description PNR login
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/booking/pnr-login [post]
// @Tags ConfirmedUser
func (bs *BookingServer) PNRLogin(ctx *gin.Context) {
	handlers.PNRLogin(ctx, bs.pb)
}

// @Summary Confirm a booking
// @Description Confirm a booking
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/booking/confirm-booking [post]
// @Tags Booking
func (bs *BookingServer) ConfirmBooking(ctx *gin.Context) {
	handlers.ConfirmBooking(ctx, bs.pb)
}

// @Summary Process online payment
// @Description Process online payment
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/booking/online-payment [post]
// @Tags Booking
func (bs *BookingServer) OnlinePayment(ctx *gin.Context) {
	handlers.OnlinePayment(ctx, bs.pb)
}

// @Summary Payment success
// @Description Payment success
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/booking/payment-success [get]
// @Tags Booking
func (bs *BookingServer) PaymentSuccess(ctx *gin.Context) {
	handlers.PaymentSuccess(ctx, bs.pb)
}

// @Summary Render payment success page
// @Description Render payment success page
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/booking/payment-success-page [get]
// @Tags Booking
func (bs *BookingServer) PaymentSuccessPage(ctx *gin.Context) {
	handlers.PaymentSuccessPage(ctx, bs.pb)
}
