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

func CreateAirlineSeat(ctx *gin.Context, client pb.AdminAirlineClient) {
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

	var req dto.AirlineSeatRequest
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

	response, err := client.RegisterAirlineSeat(newCont, &pb.AirlineSeatRequest{
		EconomySeatNo:       int32(req.EconomySeatNumber),
		BuisinesSeatNo:      int32(req.BuisinesSeatNumber),
		EconomySeatsPerRow:  int32(req.EconomySeatsPerRow),
		BuisinesSeatsPerRow: int32(req.BuisinesSeatsPerRow),
		AirlineEmail:        airlineEmails,
	})

	if err != nil {
		log.Printf("error registering airline Seat err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": "Airline Seat Created",
		"data":    response,
	})
}

func GetSeats(ctx *gin.Context, client pb.AdminAirlineClient) {
	nCtx, cancel := context.WithTimeout(ctx, time.Second*1000)
	defer cancel()

	em, _ := ctx.Get("registered_email")
	email := em.(string)
	response, err := client.FetchAllAirlineSeats(nCtx, &pb.FetchRequest{
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

func GetSeat(ctx *gin.Context, client pb.AdminAirlineClient) {
	nCtx, cancel := context.WithTimeout(ctx, time.Second*1000)
	defer cancel()

	id := ctx.Param("id")
	em, _ := ctx.Get("registered_email")
	email := em.(string)

	response, err := client.FetchAirlineSeat(nCtx, &pb.FetchRequest{
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

func DeleteSeat(ctx *gin.Context, client pb.AdminAirlineClient) {
	nCtx, cancel := context.WithTimeout(ctx, time.Second*1000)
	defer cancel()

	id := ctx.Param("id")
	em, _ := ctx.Get("registered_email")
	email := em.(string)
	response, err := client.DeleteAirlineSeat(nCtx, &pb.FetchRequest{
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
