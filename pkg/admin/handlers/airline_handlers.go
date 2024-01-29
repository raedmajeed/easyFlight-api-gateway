package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	dto "github.com/raedmajeed/api-gateway/pkg/DTO"
	pb "github.com/raedmajeed/api-gateway/pkg/admin/pb"
	"github.com/raedmajeed/api-gateway/utitlity"
)

func RegisterAirline(ctx *gin.Context, client pb.AdminAirlineClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var req dto.AirlineCompanyRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	//? Validating struct
	validate := validator.New(validator.WithRequiredStructEnabled())
	_ = validate.RegisterValidation("emailcst", utitlity.EmailValidation)
	_ = validate.RegisterValidation("phone", utitlity.PhoneNumberValidation)
	err := validate.Struct(req)
	if err != nil {
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
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
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("error in field %v, error: %v", e.Field(), e.Tag()),
			})
		}
		return
	}

	response, err := client.VerifyAirlineRegistration(cont, &pb.OTPRequest{
		Otp:   int32(req.Otp),
		Email: req.Email,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "OTP verified, Airline Creation successful. Will be sending the login credentials once verification is completed",
		"data":    response,
	})
}

func GetAllAirlines(ctx *gin.Context, client pb.AdminAirlineClient) {
	nCtx, cancel := context.WithTimeout(ctx, time.Second*1000)
	defer cancel()

	response, err := client.FetchAllAirlines(nCtx, &pb.EmptyRequest{})
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

func GetAcceptedAirlines(ctx *gin.Context, client pb.AdminAirlineClient) {
	nCtx, cancel := context.WithTimeout(ctx, time.Second*1000)
	defer cancel()

	response, err := client.GetAcceptedAirlines(nCtx, &pb.EmptyRequest{})
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

func GetRejectedAirlines(ctx *gin.Context, client pb.AdminAirlineClient) {
	nCtx, cancel := context.WithTimeout(ctx, time.Second*1000)
	defer cancel()

	response, err := client.GetRejectedAirlines(nCtx, &pb.EmptyRequest{})
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
