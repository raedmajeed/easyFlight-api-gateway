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
	"google.golang.org/grpc/metadata"
)

func CreateAirlineBaggagePolicy(ctx *gin.Context, client pb.AdminAirlineClient) {
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

	var req dto.AirlineBaggageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	//? Validating struct
	if err := validator.New().Struct(req); err != nil {
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

	airline_id := ctx.Param("airline_id")
	newCont := metadata.NewOutgoingContext(cont, metadata.Pairs("airline_id", airline_id))

	response, err := client.RegisterAirlineBaggage(newCont, &pb.AirlineBaggageRequest{
		Class:               int32(pb.Class(req.FareClass)),
		CabinAllowedWeight:  int32(req.CabinAllowedWeight),
		CabinAllowedLength:  int32(req.CabinAllowedLength),
		CabinAllowedBreadth: int32(req.CabinAllowedBreadth),
		CabinAllowedHeight:  int32(req.CabinAllowedHeight),
		HandAllowedWeight:   int32(req.HandAllowedWeight),
		HandAllowedLength:   int32(req.HandAllowedLength),
		HandAllowedBreadth:  int32(req.CabinAllowedBreadth),
		HandAllowedHeight:   int32(req.HandAllowedWeight),
		FeeForExtraKgCabin:  int32(req.FeeExtraPerKGCabin),
		FeeForExtraKgHand:   int32(req.FeeExtraPerKGHand),
		Restrictions:        req.Restrictions,
		AirlineEmail:        airlineEmails,
	})

	if err != nil {
		log.Printf("error registering airline Baggage policy err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": "Airline Baggage Policy Created",
		"data":    response,
	})
}
