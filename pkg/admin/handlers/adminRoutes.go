package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/raedmajeed/api-gateway/pkg/admin/pb"
	"google.golang.org/grpc/metadata"
)

func RegisterFlightType(ctx *gin.Context, client pb.AdminAirlineClient) {
	var reqData map[string]interface{}

	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err := ctx.BindJSON(&reqData); err != nil {
		log.Printf("Error parsing JSON request: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error": "error binding JSON request data",
		})
		return
	}

	// fmt.Println(reqData["name"])
	// metadata.NewOutgoingContext(cont, metadata.Pairs("req_data", reqData))
	newContext := context.WithValue(cont, "req_data", reqData)

	response, err := client.RegisterFlightType(newContext, &pb.FlightTypeRequest{
		Name: reqData["name"].(string),
		Tyupe: reqData["type"].(string),
	})

	if err != nil {
		log.Printf("error Calling method, err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status": http.StatusBadGateway,
			"error": "gRPC call failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Flight type retrieved successfully",
		"data": gin.H{
			"flightType": response.GetFlightType(),
		},
	})
}

func GetFlightTypes(ctx *gin.Context, client pb.AdminAirlineClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	response, err := client.FetchAllFlightTypes(cont, &pb.FlightTypeRequest{})

	if err != nil {
		log.Printf("error Calling method, err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status": http.StatusBadGateway,
			"error": "gRPC call failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Flight types retrieved successfully",
		"data": gin.H{
			"flightType": response.GetFlighTypes(),
		},
	})
}

func GetFlightType(ctx *gin.Context, client pb.AdminAirlineClient) {
	flightTypeID := ctx.Param("flight_type_id")

	// timeout := time.Second * 1000
	// cont, cancel := context.WithTimeout(ctx, timeout)
	// defer cancel()

	// newContext := context.WithValue(context.Background(), "f", flightTypeID)
	newContext := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("flight_type_id", flightTypeID))
	// ctx.Set("f", flightTypeID)
	// fmt.Println(newContext)
	response, err := client.FetchFlightType(newContext, &pb.FlightTypeRequest{})

	if err != nil {
		log.Printf("error Calling method, err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status": http.StatusBadGateway,
			"error": "gRPC call failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H {
		"status": http.StatusOK,
		"message": fmt.Sprintf("Flight type of id %s is retrieved successfuly", flightTypeID),
		"data": gin.H{
			"flightType": response.GetFlightType(),
		},
	})
}

func UpdateFlightType(ctx *gin.Context, client pb.AdminAirlineClient) {
	var reqData map[string]interface{}

	flightTypeID := ctx.Param("flight_type_id")
	if flightTypeID == "" {
		log.Println("No flight_type_id exists in url")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"status": http.StatusBadRequest,
			"error": "No flight_type_id exists in url",
		})
		return
	}

	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err := ctx.ShouldBindJSON(&reqData); err != nil {
		log.Printf("Error parsing JSON request: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error": "Invalid request data",
		})
		return
	}
	newContext := context.WithValue(cont, "req_data", reqData)
	newContext = context.WithValue(newContext, "flight_type_id", flightTypeID)
	response, err := client.UpdateFlightType(newContext, &pb.FlightTypeRequest{})

	if err != nil {
		log.Printf("error Calling method, err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status": http.StatusBadGateway,
			"error": "gRPC call failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": fmt.Sprintf("Flight type of id %s is updated successfuly", flightTypeID),
		"data": gin.H{
			"flightType": response.GetFlightType(),
		},
	})
}

func DeleteFlightType(ctx *gin.Context, client pb.AdminAirlineClient) {
	flightTypeID := ctx.Param("flight_type_id")
	if flightTypeID == "" {
		log.Println("No flight_type_id exists in url")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"status": http.StatusBadRequest,
			"error": "No flight_type_id exists in url",
		})
		return
	}

	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	newContext := context.WithValue(cont, "flight_type_id", flightTypeID)
	response, err := client.DeleteFlightType(newContext, &pb.FlightTypeRequest{})

	if err != nil {
		log.Printf("error Calling method, err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status": http.StatusBadGateway,
			"error": "gRPC call failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": fmt.Sprintf("Flight type of id %s is deleted successfuly", flightTypeID),
		"data": gin.H{
			"flightType": response.GetFlightType(),
		},
	})
}

func VerifyAirline(ctx *gin.Context, client pb.AdminAirlineClient) {
}

func GetAirline(ctx *gin.Context, client pb.AdminAirlineClient) {
}

func DeleteAirline(ctx *gin.Context, client pb.AdminAirlineClient) {
}

func GetAcceptedAirlines(ctx *gin.Context, client pb.AdminAirlineClient) {
}

func GetRejectedAirlines(ctx *gin.Context, client pb.AdminAirlineClient) {
}