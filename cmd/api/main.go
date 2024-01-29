package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	admin "github.com/raedmajeed/api-gateway/pkg/admin"
	booking "github.com/raedmajeed/api-gateway/pkg/bookingService"
	c "github.com/raedmajeed/api-gateway/pkg/config"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	conf, err := c.LoadConfigure()
	if err != nil {
		log.Printf("Error Loading Config Files, error: %v", err)
	}
	r := gin.Default()
	//r.LoadHTMLGlob(("../../html_templates/*.html"))
	admin.NewAdminRoutes(r, *conf)
	admin.NewAirlineRoutes(r, *conf)
	booking.NewBookingRoutes(r, *conf)
	go r.Run(":" + conf.PORT)

	sign := <-signalChan
	fmt.Println(sign)
}
