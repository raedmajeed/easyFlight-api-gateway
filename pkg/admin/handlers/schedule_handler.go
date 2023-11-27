package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	dto "github.com/raedmajeed/api-gateway/pkg/DTO"
	pb "github.com/raedmajeed/api-gateway/pkg/admin/pb"
	"github.com/raedmajeed/api-gateway/utitlity"
)

func CreateSchedule(ctx *gin.Context, client pb.AdminAirlineClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var req dto.ScheduleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	//? Validating struct
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("date", utitlity.Date)
	validate.RegisterValidation("time", utitlity.Time)

	if err := validate.Struct(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		for _, e := range err.(validator.ValidationErrors) {
			log.Printf("struct validation errors %v, %v", e.Field(), e.Tag())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("error in field %v, error: %v", e.Field(), e.Tag()),
			})
		}
		return
	}

	response, err := client.RegisterScheduleRequest(cont, &pb.ScheduleRequest{
		DepartureTime:    req.DepartureTime,
		ArrivalTime:      req.ArrivalTime,
		DepartureAirport: req.DepartureAirportCode,
		ArrivalAirport:   req.ArrivalAirportCode,
		DepartureDate:    req.DepartureDate,
		ArrivalDate:      req.ArrivalDate,
	})

	if err != nil {
		log.Printf("error registering schedule err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": "schedule Created",
		"data":    response,
	})
}
