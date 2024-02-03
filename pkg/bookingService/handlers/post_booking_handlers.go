package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/raedmajeed/api-gateway/pkg/bookingService/pb"
)

type JsonRequest struct {
	FlightCharted int32    `json:"flight_chart_id"`
	Seats         []string `json:"seats"`
}

func SelectSeat(ctx *gin.Context, client pb.BookingClient) {
	var seats JsonRequest
	pnr, _ := ctx.Get("user_pnr")
	pnrNo := fmt.Sprintf("%s", pnr)

	err := ctx.ShouldBindJSON(&seats)
	if err != nil {
		log.Println("error binding JSON SelectSeat() - booking-service")
	}

	var Seats []string
	for _, seat := range seats.Seats {
		s := fmt.Sprintf("%v", seat)
		Seats = append(Seats, s)
	}

	response, err := client.RegisterSelectSeat(ctx, &pb.SeatSelectRequest{
		PNR:           pnrNo,
		SeatArray:     Seats,
		FlightChartId: seats.FlightCharted,
	})
	if err != nil {
		log.Printf("error adding seats for PNR %v err: %v", pnrNo, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": "seats added succesfully",
		"data":    response,
	})
}
