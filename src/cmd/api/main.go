package main

import (
	"context"
	"log"
	"sync"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

type RouteState struct {
	sync.RWMutex
	data map[string][]interface{}
}

type App struct {
	router *gin.Engine
}

var routeState = &RouteState{
	data: make(map[string][]interface{}, 0),
}

func init() {
	log.Println("Gin cold start")

	r := gin.Default()
	app := App{
		router: r,
	}
	app.registerRoutes()

	ginLambda = ginadapter.New(app.router)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(handler)
}
