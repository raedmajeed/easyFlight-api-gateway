package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/raedmajeed/api-gateway/pkg/admin/pb"
	"github.com/raedmajeed/api-gateway/pkg/config"
)

func RegisterFlightType(ctx *gin.Context, cfg config.Configure, client pb.AdminClient) {
	ct, _ := context.WithTimeout(context.Background(), time.Minute * 10)
	
	res, err := client.TestCall(ct, &pb.TestHello{
		Message: "Test call from api gateway",
	})

	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatus(http.StatusBadGateway)
	}

	fmt.Println(res.Message)
}

func GetFlightTypes(ctx *gin.Context, cfg config.Configure, client pb.AdminClient) {
	ct, _ := context.WithTimeout(context.Background(), time.Minute * 10)

	res, err := client.TestCall(ct, &pb.TestHello{
		Message: "Test call from api gateway",
	})

	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatus(http.StatusBadGateway)
	}

	fmt.Println(res.Message)
}

func GetFlightType(ctx *gin.Context, cfg config.Configure, client pb.AdminClient) {
	ct, _ := context.WithTimeout(context.Background(), time.Minute * 10)

	res, err := client.TestCall(ct, &pb.TestHello{
		Message: "Test call from api gateway",
	})

	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatus(http.StatusBadGateway)
	}

	fmt.Println(res.Message)
}

func UpdateFlightType(ctx *gin.Context, cfg config.Configure, client pb.AdminClient) {
	ct, _ := context.WithTimeout(context.Background(), time.Minute * 10)

	res, err := client.TestCall(ct, &pb.TestHello{
		Message: "Test call from api gateway",
	})

	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatus(http.StatusBadGateway)
	}

	fmt.Println(res.Message)
}

func DeleteFlightType(ctx *gin.Context, cfg config.Configure, client pb.AdminClient) {
	ct, _ := context.WithTimeout(context.Background(), time.Minute * 10)

	res, err := client.TestCall(ct, &pb.TestHello{
		Message: "Test call from api gateway",
	})

	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatus(http.StatusBadGateway)
	}

	fmt.Println(res.Message)
}
