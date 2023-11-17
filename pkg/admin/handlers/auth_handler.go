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
	"google.golang.org/grpc/metadata"
)

func Login(ctx *gin.Context, client pb.AdminAirlineClient, role string) {
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
		Role:     role,
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

// ForgotPasswordRequest
func ForgotPasswordRequest(ctx *gin.Context, client pb.AdminAirlineClient) {
	var req dto.ForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

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

	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	response, err := client.RegisterForgotPasswordRequest(cont, &pb.ForgotPasswordRequest{
		Email: req.Email,
	})

	if err != nil {
		log.Printf("user %v not found err: %v", req.Email, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("OTP generated to email %v", req.Email),
		"data":    response,
	})
}

// * Confirm OTP Request
func ConfirmOTPRequest(ctx *gin.Context, client pb.AdminAirlineClient) {
	var req dto.OTP
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

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

	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	response, err := client.RegisterVerifyOTPRequest(cont, &pb.OTPRequest{
		Email: req.Email,
		Otp:   int32(req.Otp),
	})

	if err != nil {
		log.Printf("OTP not verified, unable to generate password for user %v err: %v", req.Email, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("OTP verification succesful for email %v", req.Email),
		"data":    response,
	})
}

// *Reset Password Request
func ConfirmPasswordReq(ctx *gin.Context, client pb.AdminAirlineClient) {
	var req dto.ConfirmPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.VarWithValue(req.ConfirmPassword, req.Password, "eqfield")
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

	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	email, check := ctx.Get("registered_email")
	if !check {
		log.Println("unable to get value from context")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintln("unable to get value from context"),
		})
	}

	emailz := fmt.Sprintf("%v", email)
	newCont := metadata.NewOutgoingContext(cont, metadata.Pairs("registered_email", emailz))
	response, err := client.RegisterConfirmPasswordRequest(newCont, &pb.ConfirmPasswordRequest{
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	})

	if err != nil {
		log.Printf("Password not reset for user err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("Password reset succesful"),
		"data":    response,
	})
}
