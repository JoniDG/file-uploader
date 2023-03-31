package job

import (
	"context"
	"file-uploader/internal/controller"
	"file-uploader/internal/defines"
	"file-uploader/internal/repository"
	"file-uploader/internal/service"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

type Jobs interface {
	Start()
}
type job struct {
}

func NewJob() Jobs {
	return &job{}
}

func InitPostgres() *sqlx.DB {
	postgresURI := fmt.Sprintf("postgres://%s:%s@%s/postgres?sslmode=disable",
		os.Getenv(defines.EnvPostgresUser),
		os.Getenv(defines.EnvPostgresPassword),
		os.Getenv(defines.EnvPostgresHost),
	)
	db, err := sqlx.Open("postgres", postgresURI)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Postgres Conected")
	return db
}

func InitRedis() *redis.Client {
	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(defines.EnvRedisHost),
		Password: os.Getenv(defines.EnvRedisPassword),
	})
	err := redisClient.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("Error ping Redis: %+v\n", err)
	}
	log.Println("Redis Connected")
	return redisClient
}

func (c *job) Start() {
	//PostgresSQL init
	db := InitPostgres()
	// Repositories init
	repo := repository.NewUserRepository(db)

	// Services init
	svc := service.NewFileService(repo)

	// Controllers init
	ctrl := controller.NewUploadController(svc)

	//Redis Client Init
	rc := InitRedis()
	ctx := context.Background()
	fmt.Printf("Queue %s running\n", defines.QueueUploadFile)
	for {
		result, err := rc.BLPop(ctx, 0, defines.QueueUploadFile).Result()
		if err != nil {
			log.Println(err)
		}
		if result != nil && len(result) == 2 {
			go ctrl.UploadFile(&result[1])
		}

	}
}
