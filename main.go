package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"pairs/src/api/routes"
)

func main() {

	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)

	defer f.Close()

	route := routes.NewRouter()

	log.Println("Data Layer Server Started")
	httpErr := http.ListenAndServe(":3341", route.Router)
	if err != nil {
		log.Println(httpErr)
	}
}
