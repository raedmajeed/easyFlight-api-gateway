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
)

func CreateAirport(ctx *gin.Context, client pb.AdminAirlineClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var req dto.AirportRequest
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

	response, err := client.RegisterAirportRequest(cont, &pb.Airport{
		AirportCode:  req.AirportCode,
		AirportName:  req.AirportName,
		City:         req.City,
		Country:      req.Country,
		Region:       req.Region,
		Latitude:     req.Latitude,
		Longitude:    req.Longitude,
		IataFcsCode:  req.IATAFCSCode,
		IcaoCode:     req.ICAOCode,
		Website:      req.Website,
		ContactEmail: req.ContactEmail,
		ContactPhone: req.ContactPhone,
	})

	if err != nil {
		log.Printf("error registering airport err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": "airport Created",
		"data":    response,
	})
}
