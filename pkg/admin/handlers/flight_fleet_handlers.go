package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	dto "github.com/raedmajeed/api-gateway/pkg/DTO"
	pb "github.com/raedmajeed/api-gateway/pkg/admin/pb"
)

func AddFleetList(ctx *gin.Context, client pb.AdminAirlineClient) {
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
	var req dto.FlightFleetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
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

	response, err := client.RegisterFlightFleets(cont, &pb.FlightFleetRequest{
		AirlineEmail:         airlineEmails,
		SeatId:               int32(req.SeatId),
		FlightTypeId:         int32(req.FlightTypeId),
		BaggagePolicyId:      int32(req.BaggagePolicyId),
		CancellationPolicyId: int32(req.CancellationPolicyId),
	})

	if err != nil {
		log.Printf("error adding flights: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": "flight fleets added",
		"data":    response,
	})
}

func GetFleetFlight(ctx *gin.Context, client pb.AdminAirlineClient) {
	nCtx, cancel := context.WithTimeout(ctx, time.Second*1000)
	defer cancel()

	id := ctx.Param("id")
	em, _ := ctx.Get("registered_email")
	email := em.(string)
	response, err := client.GetFlightFleet(nCtx, &pb.FetchRequest{
		Id:    id,
		Email: email,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "response fetched successfully",
		"data":    response,
	})
}

func GetFleetFlights(ctx *gin.Context, client pb.AdminAirlineClient) {
	nCtx, cancel := context.WithTimeout(ctx, time.Second*1000)
	defer cancel()

	em, _ := ctx.Get("registered_email")
	email := em.(string)
	response, err := client.GetFlightFleets(nCtx, &pb.FetchRequest{
		Email: email,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "response fetched successfully",
		"data":    response,
	})
}

func DeleteFleet(ctx *gin.Context, client pb.AdminAirlineClient) {
	nCtx, cancel := context.WithTimeout(ctx, time.Second*1000)
	defer cancel()

	id := ctx.Param("id")
	em, _ := ctx.Get("registered_email")
	email := em.(string)
	response, err := client.DeleteFlightFleet(nCtx, &pb.FetchRequest{
		Id:    id,
		Email: email,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "data deleted successfully",
		"data":    response,
	})
}
