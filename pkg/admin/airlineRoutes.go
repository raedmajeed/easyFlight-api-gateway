package admin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raedmajeed/api-gateway/middleware"
	handler "github.com/raedmajeed/api-gateway/pkg/admin/handlers"
	pb "github.com/raedmajeed/api-gateway/pkg/admin/pb"
	"github.com/raedmajeed/api-gateway/pkg/config"
)

type Airline struct {
	c      *gin.Engine
	cfg    *config.Configure
	client pb.AdminAirlineClient
}

func NewAirlineRoutes(c *gin.Engine, cfg config.Configure) {
	// Dialing to gRPC client admin & airline
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error Not connected with gRPC server, %v", err.Error())
	}
	airlineHandler := &Airline{
		c:      c,
		cfg:    &cfg,
		client: client,
	}

	apiVersion := c.Group("/api/v1")
	airline := apiVersion.Group("/airline")
	{
		// Register an airline
		airline.POST("/register", airlineHandler.RegisterAirline)

		// Verify airline registration
		airline.POST("/verify-registration", airlineHandler.VerifyRegistration)

		// Update airline information
		airline.PUT("/:airline_id", airlineHandler.AirlineAuthenticate, airlineHandler.UpdateAirline)

		// airline login
		airline.POST("/login", airlineHandler.AirlineLogin)

		// airline forgot password
		airline.POST("/forgot", airlineHandler.ForgotPassword)
		airline.POST("/forgot/verify", airlineHandler.VerifyAirline)
		airline.POST("/forgot/verify/reset", airlineHandler.AirlineAuthenticate, airlineHandler.ResetPassword)

		// Seats routes
		seats := airline.Group("/seats", airlineHandler.AirlineAuthenticate)
		{
			seats.POST("", airlineHandler.CreateAirlineSeat)
			seats.GET("", airlineHandler.GetSeats)
			seats.GET("/:id", airlineHandler.GetSeat)
			seats.DELETE("/:id", airlineHandler.DeleteSeat)
		}

		// Baggage Policies routes
		baggagePolicies := airline.Group("/baggage-policies", airlineHandler.AirlineAuthenticate)
		{
			baggagePolicies.POST("", airlineHandler.CreateAirlineBaggagePolicy)
			baggagePolicies.DELETE("/:id", airlineHandler.DeleteBaggagePolicy)
			baggagePolicies.GET("", airlineHandler.GetBaggagePolicies)
			baggagePolicies.GET("/:id", airlineHandler.GetBaggagePolicy)
		}

		// Cancellation Policies routes
		cancellationPolicies := airline.Group("/cancellation-policies", airlineHandler.AirlineAuthenticate)
		{
			cancellationPolicies.POST("", airlineHandler.CreateAirlineCancellationPolicy)
			cancellationPolicies.DELETE("/:id", airlineHandler.DeleteCancellationPolicy)
			cancellationPolicies.GET("", airlineHandler.GetCancellationPolicies)
			cancellationPolicies.GET("/:id", airlineHandler.GetCancellationPolicy)
		}

		// Fleet routes
		fleet := airline.Group("/fleet", airlineHandler.AirlineAuthenticate)
		{
			fleet.POST("", airlineHandler.AddFleetList)
			fleet.DELETE("/:id", airlineHandler.DeleteFleet)
			fleet.GET("/flights", airlineHandler.GetFleetFlights)
			fleet.GET("/flights/:id", airlineHandler.GetFleetFlight)
			//fleet.PATCH("/:fleet_id/maintenance", airlineHandler.UpdateFleetMaintenance)
		}

		// Flight Charts routes
		flightCharts := airline.Group("/flight-charts", airlineHandler.AirlineAuthenticate)
		{
			flightCharts.POST("", airlineHandler.CreateFlightChart)
			flightCharts.GET("/:id", airlineHandler.GetFlightChartForAirline)
		}
	}
}

func (a *Airline) AirlineAuthenticate(ctx *gin.Context) {
	email, err := middleware.ValidateToken(ctx, *a.cfg, "airline")
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

func (a *Airline) RegisterAirline(ctx *gin.Context) {
	handler.RegisterAirline(ctx, a.client)
}

func (a *Airline) VerifyRegistration(ctx *gin.Context) {
	handler.VerifyRegistration(ctx, a.client)
}

func (a *Airline) UpdateAirline(ctx *gin.Context) {
	handler.RegisterAirline(ctx, a.client)
}

// CreateAirlineSeat Below functions deals with all operations related to seats
func (a *Airline) CreateAirlineSeat(ctx *gin.Context) {
	handler.CreateAirlineSeat(ctx, a.client)
}

func (a *Airline) GetSeats(ctx *gin.Context) {
	handler.GetSeats(ctx, a.client)
}

func (a *Airline) GetSeat(ctx *gin.Context) {
	handler.GetSeat(ctx, a.client)
}

func (a *Airline) DeleteSeat(ctx *gin.Context) {
	handler.DeleteSeat(ctx, a.client)
}

// CreateAirlineBaggagePolicy Below functions deals with all operations related to seats
func (a *Airline) CreateAirlineBaggagePolicy(ctx *gin.Context) {
	handler.CreateAirlineBaggagePolicy(ctx, a.client)
}

func (a *Airline) GetBaggagePolicies(ctx *gin.Context) {
	handler.GetBaggagePolicies(ctx, a.client)
}

func (a *Airline) GetBaggagePolicy(ctx *gin.Context) {
	handler.GetBaggagePolicy(ctx, a.client)
}

func (a *Airline) DeleteBaggagePolicy(ctx *gin.Context) {
	handler.DeleteBaggagePolicy(ctx, a.client)
}

// CreateAirlineCancellationPolicy Below functions deals with all operations related to seats
func (a *Airline) CreateAirlineCancellationPolicy(ctx *gin.Context) {
	handler.CreateAirlineCancellationPolicy(ctx, a.client)
}

func (a *Airline) GetCancellationPolicies(ctx *gin.Context) {
	handler.GetCancellationPolicies(ctx, a.client)
}

func (a *Airline) GetCancellationPolicy(ctx *gin.Context) {
	handler.GetCancellationPolicy(ctx, a.client)
}

func (a *Airline) DeleteCancellationPolicy(ctx *gin.Context) {
	handler.DeleteCancellationPolicy(ctx, a.client)
}

func (a *Airline) AirlineLogin(ctx *gin.Context) {
	handler.Login(ctx, a.client, "airline")
}

func (a *Airline) ForgotPassword(ctx *gin.Context) {
	handler.ForgotPasswordRequest(ctx, a.client)
}

func (a *Airline) VerifyAirline(ctx *gin.Context) {
	handler.ConfirmOTPRequest(ctx, a.client)
}

func (a *Airline) ResetPassword(ctx *gin.Context) {
	handler.ConfirmPasswordReq(ctx, a.client)
}

func (a *Airline) CreateFlightChart(ctx *gin.Context) {
	handler.CreateFlightChart(ctx, a.client)
}

func (a *Airline) GetFlightChartForAirline(ctx *gin.Context) {
	handler.GetFlightChartForAirline(ctx, a.client)
}

func (a *Airline) GetFleetFlight(ctx *gin.Context) {
	handler.GetFleetFlight(ctx, a.client)
}

func (a *Airline) GetFleetFlights(ctx *gin.Context) {
	handler.GetFleetFlights(ctx, a.client)
}

func (a *Airline) DeleteFleet(ctx *gin.Context) {
	handler.DeleteFleet(ctx, a.client)
}

func (a *Airline) AddFleetList(ctx *gin.Context) {
	handler.AddFleetList(ctx, a.client)
}
