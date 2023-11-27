package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	dto "github.com/raedmajeed/api-gateway/pkg/DTO"
	pb "github.com/raedmajeed/api-gateway/pkg/admin/pb"
)

func CreateFlightChart(ctx *gin.Context, client pb.AdminAirlineClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	airlineEmail, ok := ctx.Get("registered_email")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errors.New("error getting value from context"),
		})
		return
	}

	airlineEmails := fmt.Sprintf("%v", airlineEmail)
	var req dto.FlightChart
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	response, err := client.RegisterFlightChart(cont, &pb.FlightChartRequest{
		AirlineEmail:  airlineEmails,
		ScheduleId:    int32(req.ScheduleID),
		FlightFleetId: int32(req.FlightFleetID),
	})

	if err != nil {
		log.Printf("error scheduling flight: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": "flight added to schedule added",
		"data":    response,
	})
}
