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
	"google.golang.org/grpc/metadata"
)

func CreateAirlineSeat(ctx *gin.Context, client pb.AdminAirlineClient) {
	body := ctx.Request.Body
	defer body.Close()

	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

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
