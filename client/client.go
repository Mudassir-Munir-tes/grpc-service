package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Mudassir-Munir-tes/grpc-service/companypb"
	"github.com/Mudassir-Munir-tes/grpc-service/models"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer cc.Close()

	client := companypb.NewDriverServiceClient(cc)

	g := gin.Default()

	g.POST("/user", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")

		user := &models.User{}

		json.NewDecoder(ctx.Request.Body).Decode(user)

		req := &companypb.UserRequest{Request: &companypb.User{Name: user.Name}}
		result, err := client.InsertUser(ctx, req)
		if err != nil {
			log.Fatalf("error123 %v", err)
		}
		// if err == nil {
		// 	ctx.JSON(http.StatusOK, gin.H{"user": result})
		// } else {
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		// }
		json.NewEncoder(ctx.Writer).Encode(result)
	})

	g.POST("/driver", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		driver := &models.Driver{}

		json.NewDecoder(ctx.Request.Body).Decode(&driver)

		req := companypb.DriverRequest{Request: &companypb.Driver{Name: driver.Name}}

		if result, err := client.InsertDriver(ctx, &req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"driver": result})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "driver not found"})
		}

	})

	g.POST("/truck", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		truck := &models.Truck{}

		json.NewDecoder(ctx.Request.Body).Decode(&truck)

		req := companypb.TruckRequest{Request: &companypb.Truck{ModelNo: truck.ModelNo, Power: truck.Power}}

		if result, err := client.InsertTruck(ctx, &req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"truck": result})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "truck not found"})
		}

	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("failed to run server %v", err)
	}
}
