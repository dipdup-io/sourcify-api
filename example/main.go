package main

import (
	"context"
	"log"
	"time"

	"github.com/dipdup-net/sourcify-api"
)

func main() {
	api := sourcify.NewAPI("https://sourcify.dev/")

	healthCtx, healthCancel := context.WithTimeout(context.Background(), time.Second)
	defer healthCancel()

	healthResponse, err := api.Health(healthCtx)
	if err != nil {
		log.Panic(err)
	}
	log.Print(healthResponse)

	getFileCtx, getFileCancel := context.WithTimeout(context.Background(), time.Second)
	defer getFileCancel()

	file, err := api.GetFile(getFileCtx, "1", "0x3A7011e7E2b32C2B52f7De1294Ff35d6ff20310F", "full_match", "metadata.json")
	if err != nil {
		log.Panic(err)
	}
	log.Println(file)
}
