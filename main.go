package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rajritwika1/codwikz/database"
	"github.com/rajritwika1/codwikz/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect() // Connect to database

	r := gin.Default()
	//routes.SetupAuthRoutes(r)
	//	routes.SetupUserRoutes(r)
	routes.SetupAuthRoutes(r)

	port := "8080"
	fmt.Println("🚀 Server running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}
