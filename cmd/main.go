package main

import (
	"JSE_API/pkg/routes"
	"JSE_API/pkg/configs/config_colly"
	"net/http"
	"log"
)


func main() {
	config_colly.Init()
	router := routes.New()
	var err error

	httpRouter := router.GetHttpRouter()

	log.Println("Lisitening on port : 8001")
	err = http.ListenAndServe(":8001", httpRouter)
	if err != nil {
		log.Fatal("Unable to start server", err)
	}
}
