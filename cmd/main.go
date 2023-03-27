package main

import (
	"file-uploader/internal/job"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	j := job.NewJob()
	j.Job()
}
