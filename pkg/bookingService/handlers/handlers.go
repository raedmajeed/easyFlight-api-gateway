package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	pb "github.com/raedmajeed/api-gateway/pkg/bookingService/pb"
	"log"
	"net/http"
	"strconv"
	"time"
)

func SearchFlight(ctx *gin.Context, client pb.BookingClient) {
	newContext, cancel := context.WithTimeout(ctx, time.Second*1000)
	defer cancel()

	classType := ctx.Query("type")
	classTypeInt, _ := strconv.Atoi(classType)
	adults := ctx.Query("adults")
	children := ctx.Query("children")
	fromAirport := ctx.Query("fromAirport")
	toAirport := ctx.Query("toAirport")
	departDate := ctx.Query("departDate")
	page := ctx.DefaultQuery("page", "1")
	returnDate := ctx.DefaultQuery("returnDate", "")

	response, err := client.RegisterSearchFlight(newContext, &pb.SearchFlightRequest{
		Type:        pb.Type(classTypeInt),
		Adults:      adults,
		Children:    children,
		FromAirport: fromAirport,
		ToAirport:   toAirport,
		DepartDate:  departDate,
		ReturnDate:  returnDate,
		Page:        page,
	})

	if err != nil {
		log.Println("flight data not fetched, err: ", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	fmt.Println(response)
}
