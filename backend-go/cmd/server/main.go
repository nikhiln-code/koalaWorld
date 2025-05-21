package main

import (
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/config"
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/logger"
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/router"
)

/**
** Main entry point for this application.
** It loads the configuration and starts the server.
**/
func main(){
	//Initialize the logger
	isProd := false // for now we are not using the prod env
	
	logger.Init(isProd)
    log := logger.SugarLog()

    log.Info("Starting application...")
	conf, err := config.LoadConfig()
	if err != nil{
		log.Fatal("Failed to load the config:", err)
	}
	log.Info("Config loaded successfully")

	r:=router.SetupRouter()
	
    port := conf.App.Port

	log.Info("Server running at http://localhost:%s\n", port)

	if err :=r.Run(":"+port); err != nil{
		log.Fatal("Failed to start Server:", err)
	}
}