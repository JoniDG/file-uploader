package router

import (
	"context"
	"file-uploader/internal/controller"
	"file-uploader/internal/defines"
	"file-uploader/internal/repository"
	"file-uploader/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/luxarts/jsend-go"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"os"
)

func New() *gin.Engine {
	r := gin.Default()

	mapRoutes(r)

	return r
}

func InitRedisClient(ctx *context.Context) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(defines.EnvRedisHost),
		Password: os.Getenv(defines.EnvRedisPassword),
	})
	err := redisClient.Ping(*ctx).Err()
	if err != nil {
		log.Fatalf("Error ping Redis: %+v\n", err)
	}
	return redisClient
}

func mapRoutes(r *gin.Engine) {
	// DB connectors, rest clients, and other stuff init
	ctx := context.Background()
	redisClient := InitRedisClient(&ctx)
	// Repositories init
	repo := repository.NewUploadRepository()

	// Services init
	svc := service.NewUploadService(repo)

	// Controllers init
	ctrl := controller.NewUploadController(svc)

	// Endpoints
	/*r.GET(defines.EndpointExample, ctrl.ExampleHandler)*/

	// Health check endpoint
	r.GET(defines.EndpointPing, healthCheck)
	log.Printf("Polling queue %s\n", defines.QueueUploadFile)
	for {
		fileName, err := redisClient.LPop(ctx, defines.QueueUploadFile).Result()
		if err != nil {
			if err.Error() == "redis: nil" {
				continue
			}
			log.Printf("Error receiving File Name: %+v\n", err)
		}
		ctrl.UploadFile(fileName)
	}
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, jsend.NewSuccess("pong"))
}
