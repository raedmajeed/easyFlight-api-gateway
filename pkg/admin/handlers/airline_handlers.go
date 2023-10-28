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

func RegisterAirline(ctx *gin.Context, client pb.AdminAirlineClient) {
	body := ctx.Request.Body
	defer body.Close()

	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var req dto.AirlineCompanyRequest
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

	response, err := client.RegisterAirline(cont, &pb.AirlineRequest{
		AirlineName:          req.AirlineName,
		CompanyAddress:       req.CompanyAddress,
		PhoneNumber:          req.PhoneNumber,
		Email:                req.Email,
		AirlineCode:          req.AirlineCode,
		AirlineLogoLink:      req.AirlineLogoLink,
		SupportDocumentsLink: req.SupportDocumentLink,
	})

	if err != nil {
		log.Printf("error registering airline err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": "Airline Creation initiated, check mail for OTP",
		"data":    response,
	})
}

func VerifyRegistration(ctx *gin.Context, client pb.AdminAirlineClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var req dto.OTP
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

	response, err := client.VerifyAirline(cont, &pb.OTPRequest{
		OTP:   int32(req.Otp),
		Email: req.Email,
	})

	if err != nil {
		log.Printf("error registering airline err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": "OTP verified, Airline Creation succesful",
		"data":    response,
	})
}

func UpdateAirline(ctx *gin.Context, client pb.AdminAirlineClient) {
}

// func VerifyAirline(ctx *gin.Context, client pb.AdminAirlineClient) {
// }

// func GetAirline(ctx *gin.Context, client pb.AdminAirlineClient) {
// }

// func DeleteAirline(ctx *gin.Context, client pb.AdminAirlineClient) {
// }

// func GetAcceptedAirlines(ctx *gin.Context, client pb.AdminAirlineClient) {
// }

// func GetRejectedAirlines(ctx *gin.Context, client pb.AdminAirlineClient) {
// }
