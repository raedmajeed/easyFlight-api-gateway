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

func SearchFlight(ctx *gin.Context, client pb.BookingClient) {
	newContext, cancel := context.WithTimeout(ctx, time.Second*2000)
	defer cancel()

	classType := ctx.Query("type")
	adults := ctx.Query("adults")
	children := ctx.Query("children")
	fromAirport := ctx.Query("fromAirport")
	toAirport := ctx.Query("toAirport")
	departDate := ctx.Query("departDate")
	page := ctx.DefaultQuery("page", "1")
	returnDate := ctx.DefaultQuery("returnDate", "")
	maxStops := ctx.DefaultQuery("maxStops", "0")

	response, err := client.RegisterSearchFlight(newContext, &pb.SearchFlightRequest{
		Type:        classType,
		Adults:      adults,
		Children:    children,
		FromAirport: fromAirport,
		ToAirport:   toAirport,
		DepartDate:  departDate,
		ReturnDate:  returnDate,
		Page:        page,
		MaxStops:    maxStops,
	})

	if err != nil {
		log.Println("flight data not fetched, err: ", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": "successful in retrieving search details",
		"data":    response,
	})
}

func SelectFlight(ctx *gin.Context, client pb.BookingClient) {
	DirectPathId := ctx.DefaultQuery("directPathId", "")
	ReturnPathId := ctx.DefaultQuery("returnPathId", "")
	refID := ctx.Query("refId")

	if DirectPathId == "" || refID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  "missing pathId or reference id, please add that co continue to booking page",
			"status": http.StatusBadRequest,
		})
		return
	}

	response, err := client.RegisterSearchSelect(ctx, &pb.SearchSelectRequest{
		Token:        refID,
		DirectPathId: DirectPathId,
		ReturnPathId: ReturnPathId,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "successful in retrieving search details",
		"data":    response,
	})
}

type TravellerDetail struct {
	Name   string `json:"name"`
	Age    string `json:"age"`
	Gender string `json:"gender"`
}

type TravellerDetails struct {
	Travellers []map[string]string `json:"travellers"`
}

func AddTravellers(ctx *gin.Context, client pb.BookingClient) {
	var travellerDetails TravellerDetails
	var td []*pb.TravellerDetails
	if err := ctx.ShouldBindJSON(&travellerDetails); err != nil {
		log.Println("unable to bind JSON, err: ", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	userEmail, ok := ctx.Get("registered_email")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errors.New("error getting value from context"),
		})
		return
	}

	userEmails := fmt.Sprintf("%v", userEmail)

	token := ctx.Query("search-token")
	for _, travellerMap := range travellerDetails.Travellers {
		td = append(td, &pb.TravellerDetails{
			Name:   travellerMap["name"],
			Age:    travellerMap["age"],
			Gender: travellerMap["gender"],
		})
	}
	response, err := client.RegisterTravellerDetails(ctx, &pb.TravellerRequest{
		TravellerDetails: td,
		Token:            token,
		Email:            userEmails,
	})

	if err != nil {
		log.Println("unable to bind JSON, err: ", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "successful in retrieving search details",
		"data":    response,
	})
}
