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
	validate.RegisterValidation("alphaspace", utitlity.AlphaSpace)
	validate.RegisterValidation("emailcst", utitlity.EmailValidation)
	validate.RegisterValidation("phone", utitlity.PhoneNumberValidation)

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

//func GetAirport(ctx *gin.Context, pb pb.AdminAirlineClient) {
//	timeLimit := time.Second * 1000
//	context, cancel := context.WithTimeout(ctx, timeLimit)
//	defer cancel()
//
//	registeredMail, ok := ctx.Get("registered_email")
//	if !ok {
//		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
//			"status": http.StatusBadGateway,
//			"error":  errors.New("unable to get the logged in email from context"),
//		})
//	}
//
//	email := registeredMail.(string)
//	var req dto.FetchAirport
//	if err := ctx.ShouldBindJSON()
//	//response, err := pb.
//}

//func UpdateAirport(ctx *gin.Context, pb *pb.AdminAirlineClient) {
//	timeLimit := time.Second * 1000
//	context, cancel := context.WithTimeout(ctx, timeLimit)
//	defer cancel()
//
//	registeredMail, ok := ctx.Get("registered_email")
//	if !ok {
//		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
//			"status": http.StatusBadGateway,
//			"error":  errors.New("unable to get the logged in email from context"),
//		})
//	}
//
//	email := registeredMail.(string)
//}
