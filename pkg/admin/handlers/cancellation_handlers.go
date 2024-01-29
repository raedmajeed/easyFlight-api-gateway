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

func CreateAirlineCancellationPolicy(ctx *gin.Context, client pb.AdminAirlineClient) {
	body := ctx.Request.Body
	defer body.Close()

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

	var req dto.AirlineCancellationRequest
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

	response, err := client.RegisterAirlineCancellation(newCont, &pb.AirlineCancellationRequest{
		Class:                           pb.Class(req.FareClass - 1), // Convert to the protobuf enum
		CancellationDeadlineBeforeHours: uint32(req.CancellationDeadlineBefore),
		CancellationPercentage:          int32(req.CancellationPercentage),
		Refundable:                      req.Refundable,
		AirlineEmail:                    airlineEmails,
	})

	if err != nil {
		log.Printf("error registering airline Cancellation err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": "Airline Cancellation Policy Created",
		"data":    response,
	})
}

func GetCancellationPolicies(ctx *gin.Context, client pb.AdminAirlineClient) {
	nCtx, cancel := context.WithTimeout(ctx, time.Second*1000)
	defer cancel()

	em, _ := ctx.Get("registered_email")
	email := em.(string)
	response, err := client.FetchAllAirlineCancellations(nCtx, &pb.FetchRequest{
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

func GetCancellationPolicy(ctx *gin.Context, client pb.AdminAirlineClient) {
	nCtx, cancel := context.WithTimeout(ctx, time.Second*1000)
	defer cancel()

	id := ctx.Param("id")
	em, _ := ctx.Get("registered_email")
	email := em.(string)
	response, err := client.FetchAirlineCancellation(nCtx, &pb.FetchRequest{
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

func DeleteCancellationPolicy(ctx *gin.Context, client pb.AdminAirlineClient) {
	nCtx, cancel := context.WithTimeout(ctx, time.Second*1000)
	defer cancel()

	id := ctx.Param("id")
	em, _ := ctx.Get("registered_email")
	email := em.(string)
	response, err := client.DeleteAirlineCancellation(nCtx, &pb.FetchRequest{
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
