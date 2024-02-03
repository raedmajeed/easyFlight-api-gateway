package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	_ "github.com/raedmajeed/api-gateway/docs"
	admin "github.com/raedmajeed/api-gateway/pkg"
	booking "github.com/raedmajeed/api-gateway/pkg"
	c "github.com/raedmajeed/api-gateway/pkg/config"
)

// @title           Flight Booking System
// @version         1.0
// @description     This is a flight booking system where users can book flights giveing the departure and arrival airport.
// @description 		Users have the option to filter flight based on max stops.
// @termsOfService  http://swagger.io/terms/

// @contact.name   EasyFlight Support
// @contact.url    http://www.easyflight.com/support
// @contact.email  support@easyflight.com

// @host      localhost:8086
// @BasePath  /api/v1
func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	conf, err := c.LoadConfigure()
	if err != nil {
		log.Printf("Error Loading Config Files, error: %v", err)
	}
	r := gin.Default()
	admin.NewAdminRoutes(r, *conf)
	admin.NewAirlineRoutes(r, *conf)
	booking.NewBookingRoutes(r, *conf)
	go r.Run(":" + conf.PORT)

	sign := <-signalChan
	fmt.Println(sign)
}
