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

func GetFlightChart(ctx *gin.Context, client pb.AdminAirlineClient) {
	timeLimit := time.Second * 1000
	newCtx, cancel := context.WithTimeout(ctx, timeLimit)
	defer cancel()

	depAirport := ctx.Query("dep")
	arrAirport := ctx.Query("arr")

	response, err := client.GetFlightChart(newCtx, &pb.GetChartRequest{
		DepAirport: depAirport,
		ArrAirport: arrAirport,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "false",
			"error":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "true",
		"message": "response fetched successfully",
		"data":    response,
	})
}

func GetFlightCharts(ctx *gin.Context, client pb.AdminAirlineClient) {
	timeLimit := time.Second * 1000
	newCtx, cancel := context.WithTimeout(ctx, timeLimit)
	defer cancel()

	response, err := client.GetFlightCharts(newCtx, &pb.EmptyRequest{})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "false",
			"error":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "true",
		"message": "response fetched successfully",
		"data":    response,
	})
}

func GetFlightChartForAirline(ctx *gin.Context, client pb.AdminAirlineClient) {
	timeLimit := time.Second * 1000
	newCtx, cancel := context.WithTimeout(ctx, timeLimit)
	defer cancel()

	em, _ := ctx.Get("registered_email")
	email := em.(string)
	id := ctx.Param("id")
	response, err := client.GetFlightChartForAirline(newCtx, &pb.FetchRequest{
		Id:    id,
		Email: email,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "false",
			"error":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "true",
		"message": "response fetched successfully",
		"data":    response,
	})
}
