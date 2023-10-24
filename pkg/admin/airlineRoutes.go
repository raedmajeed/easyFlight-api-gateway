package admin

import (
	"log"

	"github.com/gin-gonic/gin"
	handler "github.com/raedmajeed/api-gateway/pkg/admin/handlers"
	pb "github.com/raedmajeed/api-gateway/pkg/admin/pb"
	"github.com/raedmajeed/api-gateway/pkg/config"
)

type Airline struct {
	c     *gin.Engine
	cfg   *config.Configure
	client pb.AdminClient
}

func NewAirlineRoutes(c *gin.Engine, cfg config.Configure) {
	// Dialing to gRPC client admin & airline
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error Not connected with gRPC server, %v", err.Error())
	}
	airlineHandler := &Airline{
		c:     c,
		cfg:   &cfg,
		client: client,
	}

	apiVersion := c.Group("/api/v1")
	airline := apiVersion.Group("/airlines")
	{
		// Register an airline
		airline.POST("/register", airlineHandler.RegisterAirline)

		// Verify airline registration
		airline.POST("/verify-registration", airlineHandler.VerifyRegistration)

		// Update airline information
		airline.PUT("/:airline_id", airlineHandler.UpdateAirline)

		// Seats routes
		seats := airline.Group("/:airline_id/seats")
		{
			seats.POST("", airlineHandler.CreateSeat)
			seats.PUT("/:seat_id", airlineHandler.UpdateSeat)
			seats.GET("", airlineHandler.GetSeats)
			seats.DELETE("/:seat_id", airlineHandler.DeleteSeat)
		}

		// Baggage Policies routes
		baggagePolicies := airline.Group("/:airline_id/baggage-policies")
		{
			baggagePolicies.POST("", airlineHandler.CreateBaggagePolicy)
			baggagePolicies.PUT("/:policy_id", airlineHandler.UpdateBaggagePolicy)
			baggagePolicies.DELETE("/:policy_id", airlineHandler.DeleteBaggagePolicy)
			baggagePolicies.GET("", airlineHandler.GetBaggagePolicies)
		}

		// Cancellation Policies routes
		cancellationPolicies := airline.Group("/:airline_id/cancellation-policies")
		{
			cancellationPolicies.POST("", airlineHandler.CreateCancellationPolicy)
			cancellationPolicies.PUT("/:policy_id", airlineHandler.UpdateCancellationPolicy)
			cancellationPolicies.DELETE("/:policy_id", airlineHandler.DeleteCancellationPolicy)
			cancellationPolicies.GET("", airlineHandler.GetCancellationPolicies)
		}

		// Fleet routes
		fleet := airline.Group("/:airline_id/fleet")
		{
			fleet.GET("", airlineHandler.GetFleetList)
			fleet.PUT("/:fleet_id", airlineHandler.UpdateFleet)
			fleet.DELETE("/:fleet_id", airlineHandler.DeleteFleet)
			fleet.GET("/flights", airlineHandler.GetFleetFlights)
			fleet.PATCH("/:fleet_id/maintenance", airlineHandler.UpdateFleetMaintenance)
		}

		// Flight Charts routes
		flightCharts := airline.Group("/:airline_id/flight-charts")
		{
			flightCharts.POST("", airlineHandler.CreateFlightChart)
			flightCharts.PATCH("/:chart_id/status", airlineHandler.UpdateFlightChartStatus)
			flightCharts.DELETE("/:chart_id", airlineHandler.DeleteFlightChart)
			flightCharts.GET("", airlineHandler.GetFlightCharts)
		}

		// Cancellation Amount route
		airline.GET("/:airline_id/cancellation-amount", airlineHandler.GetCancellationAmount)
	}
}

func (a *Airline) registerFlightType(ctx *gin.Context) {
	handler.RegisterFlightType(ctx, *a.cfg, a.client)
}