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
	validate := validator.New(validator.WithRequiredStructEnabled())
	_ = validate.RegisterValidation("alphaspace", utitlity.AlphaSpace)
	_ = validate.RegisterValidation("emailcst", utitlity.EmailValidation)
	_ = validate.RegisterValidation("phone", utitlity.PhoneNumberValidation)

	if err := validate.Struct(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		for _, e := range err.(validator.ValidationErrors) {
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

func GetAirport(ctx *gin.Context, client pb.AdminAirlineClient) {
	timeLimit := time.Second * 1000
	newCtx, cancel := context.WithTimeout(ctx, timeLimit)
	defer cancel()

	airportCode, _ := ctx.GetQuery("airport_code")

	response, err := client.GetAirport(newCtx, &pb.AirportRequest{
		AirportCode: airportCode,
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

func GetAirports(ctx *gin.Context, client pb.AdminAirlineClient) {
	timeLimit := time.Second * 1000
	newCtx, cancel := context.WithTimeout(ctx, timeLimit)
	defer cancel()

	response, err := client.GetAirports(newCtx, &pb.EmptyRequest{})
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

func DeleteAirport(ctx *gin.Context, client pb.AdminAirlineClient) {
	timeLimit := time.Second * 1000
	newCtx, cancel := context.WithTimeout(ctx, timeLimit)
	defer cancel()

	airportCode, _ := ctx.GetQuery("airport_code")

	response, err := client.DeleteAirport(newCtx, &pb.AirportRequest{
		AirportCode: airportCode,
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
		"message": "data deleted successfully",
		"data":    response,
	})
}
