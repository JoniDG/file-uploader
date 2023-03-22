package main

import (
	"file-uploader/internal/router"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

func main() {
	r := router.New()
	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
