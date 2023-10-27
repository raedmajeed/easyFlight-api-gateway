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

	response, err := client.RegisterAirlineSeat(cont, &pb.AirlineSeatRequest{
		AirlineId:           int32(req.AirlineId),
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

func CreateAirlineBaggagePolicy(ctx *gin.Context, client pb.AdminAirlineClient) {
	body := ctx.Request.Body
	defer body.Close()

	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

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

	response, err := client.RegisterAirlineBaggage(cont, &pb.AirlineBaggageRequest{
		AirlineId:             int32(req.AirlineId),
		Class:                 pb.Class(req.FareClass),
		CabinAllowedWeight:    int32(req.CabinAllowedWeight),
		CabinAllowedDimension: int32(req.CabinAllowedDimension),
		HandAllowedWeight:     int32(req.HandAllowedWeight),
		HandAllowedDimension:  int32(req.HandAllowedDimension),
		FeeForExtraKgCabin:    int32(req.FeeExtraPerKGCabin),
		FeeForExtraKgHand:     int32(req.FeeExtraPerKGHand),
		Restrictions:          req.Restrictions,
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

func CreateAirlineCancellationPolicy(ctx *gin.Context, client pb.AdminAirlineClient) {
	body := ctx.Request.Body
	defer body.Close()

	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

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

	response, err := client.RegisterAirlineCancellation(cont, &pb.AirlineCancellationRequest{
		AirlineId:                       int32(req.AirlineId),
		Class:                           pb.Class(req.FareClass), // Convert to the protobuf enum
		CancellationDeadlineBeforeHours: uint32(req.CancellationDeadlineBefore),
		CancellationPercentage:          int32(req.CancellationPercentage),
		Refundable:                      req.Refundable,
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
