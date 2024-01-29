package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	pb "github.com/raedmajeed/api-gateway/pkg/bookingService/pb"
	"log"
	"net/http"
	"time"
)

func ConfirmBooking(ctx *gin.Context, client pb.BookingClient) {
	bookingReference := ctx.DefaultQuery("bookingRef", "")
	refId := ctx.DefaultQuery("refId", "")

	if refId == "" || bookingReference == "" {
		log.Println("reference id or booking id is empty")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errors.New("reference id or booking id is empty"),
		})
		return
	}
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	val, ok := ctx.Get("registered_email")
	if !ok {
		log.Println("email id not present in jwt token, please login again")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errors.New("email id not present in jwt token, please login again"),
		})
		return
	}

	email := fmt.Sprintf("%v", val)
	response, err := client.RegisterConfirmBooking(cont, &pb.ConfirmBookingRequest{
		Email:            email,
		Token:            refId,
		BookingReference: bookingReference,
	})

	if err != nil {
		log.Printf("unable to laod payment page %v err: %v", email, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v Make payment before 10 minutes", email),
		"data":    response,
	})
}

func OnlinePayment(ctx *gin.Context, client pb.BookingClient) {
	//headerToken := ctx.GetHeader("Authorization")
	//bearerToken := string([]byte(headerToken)[7:])

	bookingReference := ctx.DefaultQuery("bookingRef", "")
	token := ctx.DefaultQuery("searchToken", "")
	if bookingReference == "" || token == "" {
		log.Println("booking reference is not present")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errors.New("booking reference is empty"),
		})
		return
	}
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	//val, ok := ctx.Get("registered_email")
	//if !ok {
	//	log.Println("email id not present in jwt token, please login again")
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	//		"status": http.StatusBadRequest,
	//		"error":  errors.New("email id not present in jwt token, please login again"),
	//	})
	//	return
	//}
	//
	//email := fmt.Sprintf("%v", val)
	email := "raedam786@gmail.com"
	response, err := client.RegisterOnlinePayment(cont, &pb.OnlinePaymentRequest{
		Email:            email,
		BookingReference: bookingReference,
		Token:            token,
	})

	if err != nil {
		log.Printf("payment unsuccesful %v err: %v", email, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "details got successfully",
		"data":    response,
	})
}

func PaymentSuccess(ctx *gin.Context, client pb.BookingClient) {
	bookingReference := ctx.DefaultQuery("booking_reference", "")
	paymentId := ctx.DefaultQuery("payment_id", "")
	token := ctx.DefaultQuery("search_token", "")
	if bookingReference == "" {
		log.Println("booking reference is not present")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errors.New("booking reference is empty"),
		})
		return
	}
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	//val, ok := ctx.Get("registered_email")
	//if !ok {
	//	log.Println("email id not present in jwt token, please login again")
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	//		"status": http.StatusBadRequest,
	//		"error":  errors.New("email id not present in jwt token, please login again"),
	//	})
	//	return
	//}
	//
	//email := fmt.Sprintf("%v", val)
	email := "raedam786@gmail.com"
	response, err := client.ResisterPaymentConfirmed(cont, &pb.PaymentConfirmedRequest{
		Email:            email,
		BookingReference: bookingReference,
		PaymentId:        paymentId,
		Token:            token,
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
		"message": fmt.Sprintf("%v Booking Confirmed", email),
		"data":    response,
	})
}

func PaymentSuccessPage(ctx *gin.Context, client pb.BookingClient) {
	ctx.HTML(http.StatusOK, "success.html", gin.H{
		"paymentID": ctx.Query("booking_reference"),
	})
}
