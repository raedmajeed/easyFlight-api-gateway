package admin

import (
	"log"

	"github.com/gin-gonic/gin"
	handler "github.com/raedmajeed/api-gateway/pkg/admin/handlers"
	pb "github.com/raedmajeed/api-gateway/pkg/admin/pb"
	"github.com/raedmajeed/api-gateway/pkg/config"
)

type Admin struct {
	c     *gin.Engine
	cfg   *config.Configure
	client pb.AdminClient
}

func NewAdminRoutes(c *gin.Engine, cfg config.Configure) {
	// Dialing to gRPC client admin & airline
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error Not connected with gRPC server, %v", err.Error())
	}
	adminHandler := &Admin{
		c:     c,
		cfg:   &cfg,
		client: client,
	}

	apiVersion := c.Group("/api/v1")
	{
		// Group routes under /api/v1/airline
		airline := apiVersion.Group("/airline")
		{
			// Airline-specific routes
			airline.GET("/:airline_id", adminHandler.GetAirline)
			airline.DELETE("/:airline_id", adminHandler.DeleteAirline)
		}

		// Group routes under /api/v1/admin
		admin := apiVersion.Group("/admin")
		{
			// Flight Types routes
			admin.POST("/flight-types", adminHandler.RegisterFlightType)
			admin.GET("/flight-types", adminHandler.GetFlightTypes)
			admin.GET("/flight-types/:flight_type_id", adminHandler.GetFlightType)
			admin.PUT("/flight-types/:flight_type_id", adminHandler.UpdateFlightType)
			admin.DELETE("/flight-types/:flight_type_id", adminHandler.DeleteFlightType)

			// Verify Airline
			admin.POST("/verify-airline/:airline_id", adminHandler.VerifyAirline)

			// Airlines routes
			admin.GET("/airlines/accepted", adminHandler.GetAcceptedAirlines)
			admin.GET("/airlines/rejected", adminHandler.GetRejectedAirlines)

			// Airports routes
			admin.POST("/airports", adminHandler.CreateAirport)
			admin.PUT("/airports/:airport_id", adminHandler.UpdateAirport)
			admin.DELETE("/airports/:airport_id", adminHandler.DeleteAirport)
			admin.GET("/airports/:airport_id", adminHandler.GetAirport)
			admin.GET("/airports", adminHandler.GetAirports)

			// Schedules routes
			admin.POST("/schedules", adminHandler.CreateSchedule)
			admin.PUT("/schedules/:schedule_id", adminHandler.UpdateSchedule)
			admin.DELETE("/schedules/:schedule_id", adminHandler.DeleteSchedule)
			admin.GET("/schedules/:schedule_id", adminHandler.GetSchedule)
			admin.GET("/schedules", adminHandler.GetSchedules)

			// Fleet routes
			admin.GET("/fleet/:fleet_id", adminHandler.GetFleet)
			admin.GET("/fleet", adminHandler.GetFleetList)

			// Flight Charts
			admin.GET("/flight-charts/:chart_id", adminHandler.GetFlightChart)
			admin.GET("/flight-charts", adminHandler.GetFlightCharts)
		}
	}
}


func (a *Admin) registerFlightType(ctx *gin.Context) {
	handler.RegisterFlightType(ctx, *a.cfg, a.client)
}
