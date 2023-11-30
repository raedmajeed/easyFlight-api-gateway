package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	dto "github.com/raedmajeed/api-gateway/pkg/DTO"
	pb "github.com/raedmajeed/api-gateway/pkg/bookingService/pb"
	"github.com/raedmajeed/api-gateway/utitlity"
	"log"
	"net/http"
	"time"
)

func PNRLogin(ctx *gin.Context, client pb.BookingClient) {
	var req dto.PNRLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("error binding JSON PNRLogin() - booking service")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.RegisterValidation("emailcst", utitlity.EmailValidation)
	if err != nil {
		return
	}
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

	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	response, err := client.RegisterPNRLogin(cont, &pb.PNRLoginRequest{
		Email: req.Email,
		PNR:   req.PNR,
	})

	if err != nil {
		log.Printf("error logging in user %v err: %v", req.Email, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v logged in succesfully", req.Email),
		"data":    response,
	})
}
