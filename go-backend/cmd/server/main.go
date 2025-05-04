package main

import (
	"log"

	"github.com/nikhiln-code/koalaWorld/go-backend/internal/router"
)

func main(){
	r:=router.SetupRouter()
	log.Println("Server running at http://localhost:8080")

	if err :=r.Run(":8080"); err != nil{
		log.Fatal("Failed to start Server:", err)
	}
}