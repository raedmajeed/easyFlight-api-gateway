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

type Admin struct {
	cfg    *config.Configure
	client pb.AdminAirlineClient
	// redis  *redis.Client
}

func NewAdminRoutes(c *gin.Engine, cfg config.Configure) {
	// Dialing to gRPC client admin & airline
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error Not connected with gRPC server, %v", err.Error())
	}
	adminHandler := &Admin{
		cfg:    &cfg,
		client: client,
	}

	apiVersion := c.Group("/api/v1")
	{
		//* Group routes under /api/v1/airline
		// airline := apiVersion.Group("/airline")
		{
			// Airline-specific routes
			// airline.GET("/:airline_id", adminHandler.GetAirline)
			// airline.DELETE("/:airline_id", adminHandler.DeleteAirline)
		}

		//* Group routes under /api/v1/admin
		admin := apiVersion.Group("/admin")
		{
			//* Logging in
			admin.POST("/login", adminHandler.AdminLogin)

			//* Flight Types routes
			admin.POST("/flight-types", adminHandler.Authenticate, adminHandler.RegisterFlightType)
			admin.GET("/flight-types", adminHandler.Authenticate, adminHandler.GetFlightTypes)
			admin.GET("/flight-types/:flight_type_id", adminHandler.Authenticate, adminHandler.GetFlightType)
			admin.PUT("/flight-types/:flight_type_id", adminHandler.Authenticate, adminHandler.UpdateFlightType)
			admin.DELETE("/flight-types/:flight_type_id", adminHandler.Authenticate, adminHandler.DeleteFlightType)

			//* Verify Airline
			// admin.POST("/verify-airline/:airline_id", adminHandler.VerifyAirline)

			//* Airlines routes
			// admin.GET("/airlines/accepted", adminHandler.GetAcceptedAirlines)
			// admin.GET("/airlines/rejected", adminHandler.GetRejectedAirlines)

			//* Airports routes
			admin.POST("/airports", adminHandler.CreateAirport)
			// admin.PUT("/airports/:airport_id", adminHandler.UpdateAirport)
			// admin.DELETE("/airports/:airport_id", adminHandler.DeleteAirport)
			// admin.GET("/airports/:airport_id", adminHandler.GetAirport)
			// admin.GET("/airports", adminHandler.GetAirports)

			//* Schedules routes
			admin.POST("/schedules", adminHandler.CreateSchedule)
			// admin.PUT("/schedules/:schedule_id", adminHandler.UpdateSchedule)
			// admin.DELETE("/schedules/:schedule_id", adminHandler.DeleteSchedule)
			// admin.GET("/schedules/:schedule_id", adminHandler.GetSchedule)
			// admin.GET("/schedules", adminHandler.GetSchedules)

			//* Fleet routes
			// admin.GET("/fleet/:fleet_id", adminHandler.GetFleet)
			// admin.GET("/fleet", adminHandler.GetFleetList)

			//* Flight Charts
			// admin.GET("/flight-charts/:chart_id", adminHandler.GetFlightChart)
			// admin.GET("/flight-charts", adminHandler.GetFlightCharts)
		}
	}
}

func (a *Admin) Authenticate(ctx *gin.Context) {
	email, err := middleware.ValidateToken(ctx, *a.cfg, "admin")
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

func (a *Admin) RegisterFlightType(ctx *gin.Context) {
	handler.RegisterFlightType(ctx, a.client)
}

func (a *Admin) GetFlightTypes(ctx *gin.Context) {
	handler.GetFlightTypes(ctx, a.client)
}

func (a *Admin) GetFlightType(ctx *gin.Context) {
	handler.GetFlightType(ctx, a.client)
}

func (a *Admin) UpdateFlightType(ctx *gin.Context) {
	handler.UpdateFlightType(ctx, a.client)
}

func (a *Admin) DeleteFlightType(ctx *gin.Context) {
	handler.DeleteFlightType(ctx, a.client)
}

// func (a *Admin) VerifyAirline(ctx *gin.Context) {
// 	handler.VerifyAirline(ctx, a.client)
// }

// func (a *Admin) GetAirline(ctx *gin.Context) {
// 	handler.GetAirline(ctx, a.client)
// }

// func (a *Admin) DeleteAirline(ctx *gin.Context) {
// 	handler.DeleteAirline(ctx, a.client)
// }

// func (a *Admin) GetAcceptedAirlines(ctx *gin.Context) {
// 	handler.GetAcceptedAirlines(ctx, a.client)
// }

// func (a *Admin) GetRejectedAirlines(ctx *gin.Context) {
// 	handler.GetRejectedAirlines(ctx, a.client)
// }

func (a *Admin) CreateAirport(ctx *gin.Context) {
	handler.CreateAirport(ctx, a.client)
}

func (a *Admin) CreateSchedule(ctx *gin.Context) {
	handler.CreateSchedule(ctx, a.client)
}

func (a *Admin) AdminLogin(ctx *gin.Context) {
	handler.Login(ctx, a.client, "admin")
}
