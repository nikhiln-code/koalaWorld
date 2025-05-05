package main

import (
	"log"

	"github.com/nikhiln-code/koalaWorld/go-backend/internal/config"
	"github.com/nikhiln-code/koalaWorld/go-backend/internal/router"
)

func main(){
	conf, err := config.LoadConfig()
	if err != nil{
		log.Fatal("Failed to load the config:", err)
	}
	r:=router.SetupRouter()
	
    port := conf.App.Port

	log.Println("Server running at http://localhost:%s", port)

	if err :=r.Run(":"+port); err != nil{
		log.Fatal("Failed to start Server:", err)
	}
}