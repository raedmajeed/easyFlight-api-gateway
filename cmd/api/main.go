package main

import (
	"log"

	"github.com/gin-gonic/gin"
	admin "github.com/raedmajeed/api-gateway/pkg/admin"
	booking "github.com/raedmajeed/api-gateway/pkg/bookingService"
	c "github.com/raedmajeed/api-gateway/pkg/config"
)

func main() {
	conf, err := c.LoadConfigure()
	if err != nil {
		log.Printf("Error Loading Config Files, error: %v", err)
	}
	r := gin.Default()
	admin.NewAdminRoutes(r, *conf)
	admin.NewAirlineRoutes(r, *conf)
	booking.NewBookingRoutes(r, *conf)
	r.Run(":" + conf.PORT)
}
