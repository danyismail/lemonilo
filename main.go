package main

import (
	"fmt"
	"lemonilo/db"
	"lemonilo/handler"
	"lemonilo/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Lemonilo..")
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	d := db.New()
	r := router.New()
	h := handler.NewHandler(d, r)
	v1 := r.Group("/api")
	h.Register(v1)
	h.HttpErrorHandler(r)

	appPort := os.Getenv("APPLICATION_PORT")
	if "" == appPort {
		log.Fatalln("key of APPLICATION_PORT are not define.")
	}

	r.Logger.Fatal(r.Start(":" + appPort))
}
