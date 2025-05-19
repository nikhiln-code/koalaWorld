package main

import (
	"log"

	"github.com/nikhiln-code/koalaWorld/backend-go/internal/config"
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/router"
)

/**
** Main entry point for this application.
** It loads the configuration and starts the server.
**/
func main(){
	
	conf, err := config.LoadConfig()
	if err != nil{
		log.Fatal("Failed to load the config:", err)
	}

	r:=router.SetupRouter()
	
    port := conf.App.Port

	log.Printf("Server running at http://localhost:%s\n", port)

	if err :=r.Run(":"+port); err != nil{
		log.Fatal("Failed to start Server:", err)
	}
}