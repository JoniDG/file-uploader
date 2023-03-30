package main

import (
	"file-uploader/internal/job"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

func main() {
	err := godotenv.Load("./files/.env")
	if err != nil {
		log.Fatalln("Error cargando archivo .env", err)
	}
	j := job.NewJob()
	j.Start()
}
