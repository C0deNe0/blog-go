package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/C0deNe0/blog-go/handlers"
	"github.com/C0deNe0/blog-go/repository"
	"github.com/C0deNe0/blog-go/services"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var echoLambda *echoadapter.EchoLambda

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found..")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not found")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv(mongoURI)))
	if err != nil {
		log.Fatal("mongodb connection error")
	}

	db := client.Database("blogDb")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, //only our frontend domain
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT"},
	}))

	//dependecy injection
	postRepo := repository.NewPostRepository(db)
	postService := services.NewPostService(postRepo)
	handlers.NewPostHandler(e, postService)

	//wraping the echo with lambda adapter
	echoLambda = echoadapter.New(e)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(handler)
}
