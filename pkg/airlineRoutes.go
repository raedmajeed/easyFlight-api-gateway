package pkg

import (
	"log"
	"net/http"

	_ "github.com/raedmajeed/api-gateway/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/raedmajeed/api-gateway/middleware"
	a "github.com/raedmajeed/api-gateway/pkg/admin"
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
	client, err := a.ClientDial(cfg)
	if err != nil {
		log.Fatalf("error Not connected with gRPC server, %v", err.Error())
	}
	airlineHandler := &Airline{
		c:      c,
		cfg:    &cfg,
		client: client,
	}

	c.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
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

// @Summary Register an airline
// @Description Register an airline
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/register [post]
// @Tags Airline
func (a *Airline) RegisterAirline(ctx *gin.Context) {
	handler.RegisterAirline(ctx, a.client)
}

// @Summary Verify airline registration
// @Description Verify airline registration
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/verify-registration [post]
// @Tags Airline
func (a *Airline) VerifyRegistration(ctx *gin.Context) {
	handler.VerifyRegistration(ctx, a.client)
}

// @Summary Update airline information
// @Description Update airline information
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/{airline_id} [put]
// @Tags Airline
func (a *Airline) UpdateAirline(ctx *gin.Context) {
	handler.RegisterAirline(ctx, a.client)
}

// @Summary Create airline seat
// @Description Create airline seat
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/seats [post]
// @Tags AirlineSeats
func (a *Airline) CreateAirlineSeat(ctx *gin.Context) {
	handler.CreateAirlineSeat(ctx, a.client)
}

// @Summary Get all seats
// @Description Get all seats
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/seats [get]
// @Tags AirlineSeats
func (a *Airline) GetSeats(ctx *gin.Context) {
	handler.GetSeats(ctx, a.client)
}

// @Summary Get a seat by ID
// @Description Get a seat by ID
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/seats/{id} [get]
// @Tags AirlineSeats
func (a *Airline) GetSeat(ctx *gin.Context) {
	handler.GetSeat(ctx, a.client)
}

// @Summary Delete a seat by ID
// @Description Delete a seat by ID
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/seats/{id} [delete]
// @Tags AirlineSeats
func (a *Airline) DeleteSeat(ctx *gin.Context) {
	handler.DeleteSeat(ctx, a.client)
}

// @Summary Create airline baggage policy
// @Description Create airline baggage policy
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/baggage-policies [post]
// @Tags AirlineBaggage
func (a *Airline) CreateAirlineBaggagePolicy(ctx *gin.Context) {
	handler.CreateAirlineBaggagePolicy(ctx, a.client)
}

// @Summary Get all baggage policies
// @Description Get all baggage policies
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/baggage-policies [get]
// @Tags AirlineBaggage
func (a *Airline) GetBaggagePolicies(ctx *gin.Context) {
	handler.GetBaggagePolicies(ctx, a.client)
}

// @Summary Get a baggage policy by ID
// @Description Get a baggage policy by ID
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/baggage-policies/{id} [get]
// @Tags AirlineBaggage
func (a *Airline) GetBaggagePolicy(ctx *gin.Context) {
	handler.GetBaggagePolicy(ctx, a.client)
}

// @Summary Delete a baggage policy by ID
// @Description Delete a baggage policy by ID
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/baggage-policies/{id} [delete]
// @Tags AirlineBaggage
func (a *Airline) DeleteBaggagePolicy(ctx *gin.Context) {
	handler.DeleteBaggagePolicy(ctx, a.client)
}

// @Summary Create airline cancellation policy
// @Description Create airline cancellation policy
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/cancellation-policies [post]
// @Tags AirlineCancellation
func (a *Airline) CreateAirlineCancellationPolicy(ctx *gin.Context) {
	handler.CreateAirlineCancellationPolicy(ctx, a.client)
}

// @Summary Get all cancellation policies
// @Description Get all cancellation policies
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/cancellation-policies [get]
// @Tags AirlineCancellation
func (a *Airline) GetCancellationPolicies(ctx *gin.Context) {
	handler.GetCancellationPolicies(ctx, a.client)
}

// @Summary Get a cancellation policy by ID
// @Description Get a cancellation policy by ID
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/cancellation-policies/{id} [get]
// @Tags AirlineCancellation
func (a *Airline) GetCancellationPolicy(ctx *gin.Context) {
	handler.GetCancellationPolicy(ctx, a.client)
}

// @Summary Delete a cancellation policy by ID
// @Description Delete a cancellation policy by ID
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/cancellation-policies/{id} [delete]
// @Tags AirlineCancellation
func (a *Airline) DeleteCancellationPolicy(ctx *gin.Context) {
	handler.DeleteCancellationPolicy(ctx, a.client)
}

// @Summary Airline login
// @Description Airline login
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/login [post]
// @Tags Airline
func (a *Airline) AirlineLogin(ctx *gin.Context) {
	handler.Login(ctx, a.client, "airline")
}

// @Summary Request password reset
// @Description Request password reset
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/forgot-password [post]
// @Tags Airline
func (a *Airline) ForgotPassword(ctx *gin.Context) {
	handler.ForgotPasswordRequest(ctx, a.client)
}

// @Summary Verify airline using OTP
// @Description Verify airline using OTP
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/verify [post]
// @Tags Airline
func (a *Airline) VerifyAirline(ctx *gin.Context) {
	handler.ConfirmOTPRequest(ctx, a.client)
}

// @Summary Reset password
// @Description Reset password
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/reset-password [post]
// @Tags Airline
func (a *Airline) ResetPassword(ctx *gin.Context) {
	handler.ConfirmPasswordReq(ctx, a.client)
}

// @Summary Create flight chart
// @Description Create flight chart
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/flight-chart [post]
// @Tags FlightChart
func (a *Airline) CreateFlightChart(ctx *gin.Context) {
	handler.CreateFlightChart(ctx, a.client)
}

// @Summary Get flight chart for airline
// @Description Get flight chart for airline
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/flight-chart [get]
// @Tags FlightChart
func (a *Airline) GetFlightChartForAirline(ctx *gin.Context) {
	handler.GetFlightChartForAirline(ctx, a.client)
}

// @Summary Get fleet flight
// @Description Get fleet flight
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/fleet-flight [get]
// @Tags FlightFleet
func (a *Airline) GetFleetFlight(ctx *gin.Context) {
	handler.GetFleetFlight(ctx, a.client)
}

// @Summary Get all fleet flights
// @Description Get all fleet flights
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/fleet-flights [get]
// @Tags FlightFleet
func (a *Airline) GetFleetFlights(ctx *gin.Context) {
	handler.GetFleetFlights(ctx, a.client)
}

// @Summary Delete fleet by ID
// @Description Delete fleet by ID
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/fleet/{id} [delete]
// @Tags FlightFleet
func (a *Airline) DeleteFleet(ctx *gin.Context) {
	handler.DeleteFleet(ctx, a.client)
}

// @Summary Add fleet list
// @Description Add fleet list
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/airline/fleet [post]
// @Tags FlightFleet
func (a *Airline) AddFleetList(ctx *gin.Context) {
	handler.AddFleetList(ctx, a.client)
}
