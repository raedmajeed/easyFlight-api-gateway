package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/raedmajeed/api-gateway/pkg/admin/pb"
	"google.golang.org/grpc/metadata"
)

func VerifyAirline(ctx *gin.Context, client pb.AdminAirlineClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	go func() {

		<-ctx.Request.Context().Done()
		// after cancel
	}()

	airlineId := ctx.Param("airline_id")
	newCont := metadata.NewOutgoingContext(cont, metadata.Pairs("airline_id", airlineId))
	response, err := client.AdminVerifyAirline(newCont, &pb.EmptyRequest{}) //! CHANGE TO MESSAGE, ADD RETRY

	if err != nil {
		log.Printf("error verifying airline err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": "Airline Verification completed, check mail for temporary password and mail",
		"data":    response,
	})
}
