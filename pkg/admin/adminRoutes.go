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
		//* Group routes under /api/v1/admin
		admin := apiVersion.Group("/admin")
		{
			//* Logging in
			admin.POST("/login", adminHandler.AdminLogin)

			//* Flight Types routes
			admin.POST("/flight-types", adminHandler.AdminAuthenticate, adminHandler.RegisterFlightType)
			admin.GET("/flight-types", adminHandler.AdminAuthenticate, adminHandler.GetFlightTypes)
			admin.GET("/flight-types/:flight_type_id", adminHandler.AdminAuthenticate, adminHandler.GetFlightType)
			admin.PUT("/flight-types/:flight_type_id", adminHandler.AdminAuthenticate, adminHandler.UpdateFlightType)
			admin.DELETE("/flight-types/:flight_type_id", adminHandler.AdminAuthenticate, adminHandler.DeleteFlightType)

			//* Verify Airline
			admin.PATCH("/verify-airline/:airline_id", adminHandler.AdminAuthenticate, adminHandler.VerifyAirline)
			//! block airline here

			//* Airlines routes
			airline := admin.Group("/airlines", adminHandler.AdminAuthenticate)
			{
				// Airline-specific routes
				airline.GET("/", adminHandler.GetAllAirlines)
				airline.GET("/accepted", adminHandler.GetAcceptedAirlines)
				airline.GET("/rejected", adminHandler.GetRejectedAirlines)

			}
			//* Airports routes
			airport := admin.Group("/airports", adminHandler.AdminAuthenticate)
			{
				airport.POST("/", adminHandler.CreateAirport)
				airport.DELETE("/", adminHandler.DeleteAirport)
				airport.GET("/airport", adminHandler.GetAirport)
				airport.GET("/", adminHandler.GetAirports)
			}

			//* Schedules routes
			schedules := admin.Group("/schedules", adminHandler.AdminAuthenticate)
			{
				schedules.POST("/", adminHandler.CreateSchedule)
				schedules.GET("/", adminHandler.GetSchedules)
			}

			//* Flight Charts
			flightCharts := admin.Group("/flight-charts", adminHandler.AdminAuthenticate)
			{
				flightCharts.GET("/flight", adminHandler.GetFlightChart)
				flightCharts.GET("/", adminHandler.GetFlightCharts)
			}
		}
	}
}

func (a *Admin) AdminAuthenticate(ctx *gin.Context) {
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

func (a *Admin) VerifyAirline(ctx *gin.Context) {
	handler.VerifyAirline(ctx, a.client)
}

func (a *Admin) GetAllAirlines(ctx *gin.Context) {
	handler.GetAllAirlines(ctx, a.client)
}

func (a *Admin) GetAcceptedAirlines(ctx *gin.Context) {
	handler.GetAcceptedAirlines(ctx, a.client)
}

func (a *Admin) GetRejectedAirlines(ctx *gin.Context) {
	handler.GetRejectedAirlines(ctx, a.client)
}

// CreateAirport below functions helps to create airport
func (a *Admin) CreateAirport(ctx *gin.Context) {
	handler.CreateAirport(ctx, a.client)
}

// GetAirport below functions helps to get an airport
func (a *Admin) GetAirport(ctx *gin.Context) {
	handler.GetAirport(ctx, a.client)
}

// DeleteAirport GetAirport below functions helps to delete an airport
func (a *Admin) DeleteAirport(ctx *gin.Context) {
	handler.DeleteAirport(ctx, a.client)
}

func (a *Admin) GetAirports(ctx *gin.Context) {
	handler.GetAirports(ctx, a.client)
}

func (a *Admin) CreateSchedule(ctx *gin.Context) {
	handler.CreateSchedule(ctx, a.client)
}

// GetSchedules GetAirports below functions helps to get all schedules
func (a *Admin) GetSchedules(ctx *gin.Context) {
	handler.GetSchedules(ctx, a.client)
}

func (a *Admin) GetFlightCharts(ctx *gin.Context) {
	handler.GetFlightCharts(ctx, a.client)
}

func (a *Admin) GetFlightChart(ctx *gin.Context) {
	handler.GetFlightChart(ctx, a.client)
}

func (a *Admin) AdminLogin(ctx *gin.Context) {
	handler.Login(ctx, a.client, "admin")
}
