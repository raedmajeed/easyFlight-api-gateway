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
		seats := airline.Group("/seats")
		{
			seats.POST("", airlineHandler.AirlineAuthenticate, airlineHandler.CreateAirlineSeat)
			//seats.PUT("/:seat_id", airlineHandler.AirlineAuthenticate, airlineHandler.UpdateSeat)
			//seats.GET("", airlineHandler.AirlineAuthenticate, airlineHandler.GetSeats)
			//seats.DELETE("/:seat_id", airlineHandler.AirlineAuthenticate, airlineHandler.DeleteSeat)
		}

		// Baggage Policies routes
		baggagePolicies := airline.Group("/baggage-policies")
		{
			baggagePolicies.POST("", airlineHandler.AirlineAuthenticate, airlineHandler.CreateAirlineBaggagePolicy)
			//baggagePolicies.PUT("/:policy_id", airlineHandler.AirlineAuthenticate, airlineHandler.UpdateBaggagePolicy)
			//baggagePolicies.DELETE("/:policy_id", airlineHandler.AirlineAuthenticate, airlineHandler.DeleteBaggagePolicy)
			//baggagePolicies.GET("", airlineHandler.AirlineAuthenticate, airlineHandler.GetBaggagePolicies)
		}

		// Cancellation Policies routes
		cancellationPolicies := airline.Group("/cancellation-policies")
		{
			cancellationPolicies.POST("", airlineHandler.AirlineAuthenticate, airlineHandler.CreateAirlineCancellationPolicy)
			//cancellationPolicies.PUT("/:policy_id", airlineHandler.AirlineAuthenticate, airlineHandler.UpdateCancellationPolicy)
			//cancellationPolicies.DELETE("/:policy_id", airlineHandler.AirlineAuthenticate, airlineHandler.DeleteCancellationPolicy)
			//cancellationPolicies.GET("", airlineHandler.GetCancellationPolicies)
		}

		// Fleet routes
		fleet := airline.Group("/fleet")
		// {
		fleet.POST("", airlineHandler.AirlineAuthenticate, airlineHandler.AddFleetList)
		//fleet.PUT("/:fleet_id", airlineHandler.UpdateFleet)
		//fleet.DELETE("/:fleet_id", airlineHandler.DeleteFleet)
		//fleet.GET("/flights", airlineHandler.GetFleetFlights)
		//fleet.PATCH("/:fleet_id/maintenance", airlineHandler.UpdateFleetMaintenance)
		// }

		// Flight Charts routes
		flightCharts := airline.Group("/flight-charts")
		{
			flightCharts.POST("", airlineHandler.AirlineAuthenticate, airlineHandler.CreateFlightChart)
			//flightCharts.GET("", airlineHandler.GetFlightCharts)
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

//func (a *Airline) GetSeats(ctx *gin.Context) {
//	handler.GetSeats(ctx, a.client)
//}
//
//func (a *Airline) UpdateSeat(ctx *gin.Context) {
//	handler.UpdateSeat(ctx, a.client)
//}
//
//func (a *Airline) DeleteSeat(ctx *gin.Context) {
//	handler.DeleteSeat(ctx, a.client)
//}

// CreateAirlineBaggagePolicy Below functions deals with all operations related to seats
func (a *Airline) CreateAirlineBaggagePolicy(ctx *gin.Context) {
	handler.CreateAirlineBaggagePolicy(ctx, a.client)
}

//func (a *Airline) GetBaggagePolicies(ctx *gin.Context) {
//	handler.GetBaggagePolicies(ctx, a.client)
//}
//
//func (a *Airline) UpdateBaggagePolicy(ctx *gin.Context) {
//	handler.UpdateBaggagePolicy(ctx, a.client)
//}
//
//func (a *Airline) DeleteBaggagePolicy(ctx *gin.Context) {
//	handler.DeleteBaggagePolicy(ctx, a.client)
//}

// CreateAirlineCancellationPolicy Below functions deals with all operations related to seats
func (a *Airline) CreateAirlineCancellationPolicy(ctx *gin.Context) {
	handler.CreateAirlineCancellationPolicy(ctx, a.client)
}

//func (a *Airline) GetCancellationPolicies(ctx *gin.Context) {
//	handler.GetCancellationPolicies(ctx, a.client)
//}
//
//func (a *Airline) UpdateCancellationPolicy(ctx *gin.Context) {
//	handler.UpdateCancellationPolicy(ctx, a.client)
//}
//
//func (a *Airline) DeleteCancellationPolicy(ctx *gin.Context) {
//	handler.DeleteCancellationPolicy(ctx, a.client)
//}

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

func (a *Airline) AddFleetList(ctx *gin.Context) {
	handler.AddFleetList(ctx, a.client)
}

func (a *Airline) CreateFlightChart(ctx *gin.Context) {
	handler.CreateFlightChart(ctx, a.client)
}
