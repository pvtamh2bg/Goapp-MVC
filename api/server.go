// Server run api and seeding data

package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pvtamh2bg/Golangwebapp/api/controllers"
	"github.com/pvtamh2bg/Golangwebapp/api/seed"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load() // Load file .env root directory
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	// initialize server connect database
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	//... excute seeding data
	seed.Load(server.DB)
	// Run server api controller
	server.Run(":8080")

}