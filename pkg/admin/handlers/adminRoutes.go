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
	"google.golang.org/grpc/metadata"
)

func RegisterFlightType(ctx *gin.Context, client pb.AdminAirlineClient) {
	ft := &dto.FlightTypeRequest{}
	// timeout := time.Second * 1000
	// cont, cancel := context.WithTimeout(ctx, timeout)
	// defer cancel()

	if err := ctx.BindJSON(ft); err != nil {
		log.Printf("Error parsing JSON request: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "error binding JSON request data",
		})
		return
	}

	//? Validating struct
	if err := validator.New().Struct(ft); err != nil {
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

	// response, err := client.RegisterFlightType(cont, &pb.FlightTypeRequest{
	// 	Type:                pb.FlightTypeEnum(ft.Type),
	// 	FlightModel:         ft.FlightModel,
	// 	Description:         ft.Description,
	// 	ManufacturerName:    ft.ManufacturerName,
	// 	ManufacturerCountry: ft.ManufacturerCountry,
	// 	MaxDistance:         ft.MaxDistance,
	// 	CruiseSpeed:         ft.CruiseSpeed,
	// })

	// if err != nil {
	// 	log.Printf("error Calling method, err: %v", err.Error())
	// 	ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
	// 		"status": http.StatusBadGateway,
	// 		"error":  "gRPC call failed",
	// 	})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"status":  http.StatusOK,
	// 	"message": fmt.Sprintf("Flight type of id %s is created successfuly", response.FlightType.FlightModel),
	// 	"data":    response,
	// })
}

func GetFlightTypes(ctx *gin.Context, client pb.AdminAirlineClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	response, err := client.FetchAllFlightTypes(cont, &pb.EmptyRequest{})

	if err != nil {
		log.Printf("error Calling method, err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status": http.StatusBadGateway,
			"error":  "gRPC call failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Flight types retrieved successfully",
		"data":    response,
	})
}

func GetFlightType(ctx *gin.Context, client pb.AdminAirlineClient) {
	flightTypeID := ctx.Param("flight_type_id")

	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	response, err := client.FetchFlightType(cont, &pb.IDRequest{})

	if err != nil {
		log.Printf("error Calling method, err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status": http.StatusBadGateway,
			"error":  "gRPC call failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Flight type of id %s is retrieved successfuly", flightTypeID),
		"data":    response,
	})
}

func UpdateFlightType(ctx *gin.Context, client pb.AdminAirlineClient) {
	ft := dto.FlightTypeRequest{}
	flightTypeID := ctx.Param("flight_type_id")
	if flightTypeID == "" {
		log.Println("No flight_type_id exists in url")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "No flight_type_id exists in url",
		})
		return
	}

	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err := ctx.ShouldBindJSON(&ft); err != nil {
		log.Printf("Error parsing JSON request: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Error parsing JSON request",
		})
		return
	}

	//? Validating struct
	if err := validator.New().Struct(ft); err != nil {
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

	newCont := metadata.NewOutgoingContext(cont, metadata.Pairs("flight_type_id", flightTypeID))
	response, err := client.UpdateFlightType(newCont, &pb.FlightTypeRequest{
		Type:                pb.FlightTypeEnum(ft.Type),
		FlightModel:         ft.FlightModel,
		Description:         ft.Description,
		ManufacturerName:    ft.ManufacturerName,
		ManufacturerCountry: ft.ManufacturerCountry,
		MaxDistance:         ft.MaxDistance,
		CruiseSpeed:         ft.CruiseSpeed,
	})

	if err != nil {
		log.Printf("error Calling method, err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status": http.StatusBadGateway,
			"error":  "gRPC call failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Flight type of id %s is updated successfuly", flightTypeID),
		"data":    response,
	})
}

func DeleteFlightType(ctx *gin.Context, client pb.AdminAirlineClient) {
	flightTypeID := ctx.Param("flight_type_id")
	if flightTypeID == "" {
		log.Println("No flight_type_id exists in url")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "No flight_type_id exists in url",
		})
		return
	}

	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	response, err := client.DeleteFlightType(cont, &pb.IDRequest{})

	if err != nil {
		log.Printf("error Calling method, err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status": http.StatusBadGateway,
			"error":  "gRPC call failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Flight type of id %s is deleted successfuly", flightTypeID),
		"data": gin.H{
			"flightType": response.GetFlightType(),
		},
	})
}

// func VerifyAirline(ctx *gin.Context, client pb.AdminAirlineClient) {
// }

// func GetAirline(ctx *gin.Context, client pb.AdminAirlineClient) {
// }

// func DeleteAirline(ctx *gin.Context, client pb.AdminAirlineClient) {
// }

// func GetAcceptedAirlines(ctx *gin.Context, client pb.AdminAirlineClient) {
// }

// func GetRejectedAirlines(ctx *gin.Context, client pb.AdminAirlineClient) {
// }
