package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/PontakornDev/ginAuth/config"
	"github.com/PontakornDev/ginAuth/repositories"
	"github.com/PontakornDev/ginAuth/router"
)

func main() {
	err := config.LoadENV()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	error := repositories.ConnectDB()
	if error != nil {
		panic("failed to connect database")
	}

	route := router.SetUpRouter()

	router.SetUpRouteGroup(route)
	migrate := flag.Bool("migrate", false, "")
	if *migrate {
		repositories.MigrationDB()
	}
	fmt.Printf("Server Listen PORT : %s\n", os.Getenv("PORT"))

	if err := route.Run(":" + os.Getenv("PORT")); err != nil {
		fmt.Println(err)
	}
}
