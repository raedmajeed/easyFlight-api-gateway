package pkg

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/raedmajeed/api-gateway/docs"
	"github.com/raedmajeed/api-gateway/middleware"

	a "github.com/raedmajeed/api-gateway/pkg/admin"
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
	client, err := a.ClientDial(cfg)
	if err != nil {
		log.Fatalf("error Not connected with gRPC server, %v", err.Error())
	}
	adminHandler := &Admin{
		cfg:    &cfg,
		client: client,
	}

	apiVersion := c.Group("/api/v2")
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

// @Summary SignUp for users
// @ID User SignUp
// @Tags User
// @Produce json
// @Router /user/signup [post]
// UserSignup handles the user signup
// @Tags FlightTypes
func (a *Admin) RegisterFlightType(ctx *gin.Context) {
	handler.RegisterFlightType(ctx, a.client)
}

// @Summary Get all flight types
// @Description Get all flight types
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/admin/flight-types [get]
// @Tags FlightTypes
func (a *Admin) GetFlightTypes(ctx *gin.Context) {
	handler.GetFlightTypes(ctx, a.client)
}

// @Summary Get a flight type by ID
// @Description Get a flight type by ID
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/admin/flight-types/{flight_type_id} [get]
// @Tags FlightTypes
func (a *Admin) GetFlightType(ctx *gin.Context) {
	handler.GetFlightType(ctx, a.client)
}

// @Summary Update a flight type by ID
// @Description Update a flight type by ID
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/admin/flight-types/{flight_type_id} [put]
// @Tags FlightTypes
func (a *Admin) UpdateFlightType(ctx *gin.Context) {
	handler.UpdateFlightType(ctx, a.client)
}

// @Summary Delete a flight type by ID
// @Description Delete a flight type by ID
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/admin/flight-types/{flight_type_id} [delete]
func (a *Admin) DeleteFlightType(ctx *gin.Context) {
	handler.DeleteFlightType(ctx, a.client)
}

// @Summary Verify airline
// @Description Verify airline
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/admin/verify-airline/{airline_id} [patch]
// @Tags Airline
func (a *Admin) VerifyAirline(ctx *gin.Context) {
	handler.VerifyAirline(ctx, a.client)
}

// @Summary Get all airlines
// @Description Get all airlines
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/admin/airlines [get]
// @Tags Airline
func (a *Admin) GetAllAirlines(ctx *gin.Context) {
	handler.GetAllAirlines(ctx, a.client)
}

// @Summary Get accepted airlines
// @Description Get accepted airlines
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/admin/airlines/accepted [get]
// @Tags Airline
func (a *Admin) GetAcceptedAirlines(ctx *gin.Context) {
	handler.GetAcceptedAirlines(ctx, a.client)
}

// @Summary Get rejected airlines
// @Description Get rejected airlines
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/admin/airlines/rejected [get]
// @Tags Airline
func (a *Admin) GetRejectedAirlines(ctx *gin.Context) {
	handler.GetRejectedAirlines(ctx, a.client)
}

// @Summary Create airport
// @Description Create airport
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/admin/airports [post]
// @Tags Airport
func (a *Admin) CreateAirport(ctx *gin.Context) {
	handler.CreateAirport(ctx, a.client)
}

// @Summary Get airport by ID
// @Description Get airport by ID
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/admin/airports/{airport_id} [get]
// @Tags Airport
func (a *Admin) GetAirport(ctx *gin.Context) {
	handler.GetAirport(ctx, a.client)
}

// @Summary Delete airport by ID
// @Description Delete airport by ID
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/admin/airports/{airport_id} [delete]
// @Tags Airport
func (a *Admin) DeleteAirport(ctx *gin.Context) {
	handler.DeleteAirport(ctx, a.client)
}

// @Summary Get all airports
// @Description Get all airports
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Router /api/v1/admin/airports [get]
// @Tags Airport
func (a *Admin) GetAirports(ctx *gin.Context) {
	handler.GetAirports(ctx, a.client)
}

// @Summary Create schedule
// @Description Create schedule
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/admin/schedules [post]
// @Tags Schedules
func (a *Admin) CreateSchedule(ctx *gin.Context) {
	handler.CreateSchedule(ctx, a.client)
}

// @Summary Get all schedules
// @Description Get all schedules
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/admin/schedules [get]
// @Tags Schedules
func (a *Admin) GetSchedules(ctx *gin.Context) {
	handler.GetSchedules(ctx, a.client)
}

// @Summary Get all flight charts
// @Description Get all flight charts
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/admin/flight-charts [get]
// @Tags FlightCharts
func (a *Admin) GetFlightCharts(ctx *gin.Context) {
	handler.GetFlightCharts(ctx, a.client)
}

// @Summary Get a flight chart by ID
// @Description Get a flight chart by ID
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/admin/flight-charts/{flight_chart_id} [get]
// @Tags FlightCharts
func (a *Admin) GetFlightChart(ctx *gin.Context) {
	handler.GetFlightChart(ctx, a.client)
}

// @Summary Admin login
// @Description Admin login
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /api/v1/admin/login [post]
// @Tags Admin
func (a *Admin) AdminLogin(ctx *gin.Context) {
	handler.Login(ctx, a.client, "admin")
}
