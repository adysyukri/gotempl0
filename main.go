package main

import (
	"log"

	"github.com/adysyukri/gotempl0/handlers"
)

func main() {
	handler := handlers.NewHandler("./static")

	log.Fatalln(handler.Listen(":3000"))
}
