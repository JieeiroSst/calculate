package main

import (
	"fmt"
	"github.com/JIeeiroSst/go-app/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
)

func main(){
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Println("client can't connect grpc server")
	}else{
		log.Println("client connect grpc server")
	}

	client := proto.NewServiceClient(conn)
	g := gin.Default()

	g.GET("/add", func(context *gin.Context) {
		a, _ := strconv.Atoi(context.Query("a"))
		b, _ := strconv.Atoi(context.Query("b"))

		req:=&proto.Request{
			A: int64(a),
			B: int64(b),
		}
		if response,err:=client.Add(context,req);err==nil{
			context.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		}else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/multiply", func(context *gin.Context) {
		a, _ := strconv.Atoi(context.Query("a"))
		b, _ := strconv.Atoi(context.Query("b"))

		req:=&proto.Request{
			A: int64(a),
			B: int64(b),
		}
		if response,err:=client.Multiply(context,req);err==nil{
			context.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		}else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/minus", func(context *gin.Context) {
		a, _ := strconv.Atoi(context.Query("a"))
		b, _ := strconv.Atoi(context.Query("b"))

		req:=&proto.Request{
			A: int64(a),
			B: int64(b),
		}
		if response,err:=client.Minus(context,req);err==nil{
			context.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		}else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/division", func(context *gin.Context) {
		a, _ := strconv.Atoi(context.Query("a"))
		b, _ := strconv.Atoi(context.Query("b"))

		req:=&proto.Request{
			A: int64(a),
			B: int64(b),
		}
		if response,err:=client.Division(context,req);err==nil{
			context.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		}else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}