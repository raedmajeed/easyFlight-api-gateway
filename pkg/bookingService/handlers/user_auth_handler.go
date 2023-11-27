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

func UserLogin(ctx *gin.Context, client pb.BookingClient) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("error binding JSON")
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

	response, err := client.RegisterLoginRequest(cont, &pb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
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

func UserRegister(ctx *gin.Context, client pb.BookingClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var req dto.UserData
	if err := ctx.BindJSON(&req); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	//? Validating struct
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("emailcst", utitlity.EmailValidation)
	validate.RegisterValidation("phone", utitlity.PhoneNumberValidation)
	validate.RegisterValidation("alphaspace", utitlity.AlphaSpace)
	err := validate.Struct(req)
	if err != nil {
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

	response, err := client.RegisterUser(cont, &pb.UserRequest{
		PhoneNumber: req.Phone,
		Email:       req.Email,
		Password:    req.Password,
		UserName:    req.Name,
	})

	if err != nil {
		log.Printf("error registering user err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": "user Creation initiated, check mail for OTP",
		"data":    response,
	})
}

func VerifyRegistration2(ctx *gin.Context, client pb.BookingClient) {
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

	response, err := client.VerifyUser(cont, &pb.OTPRequest{
		Otp:   int32(req.Otp),
		Email: req.Email,
	})

	if err != nil {
		log.Printf("error registering user err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": "OTP verified, user Creation successful.",
		"data":    response,
	})
}
