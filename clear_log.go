package main

import (
	"log"
	"log-claer/clear"
)

func main() {
	err := clear.DeleteLog("/log")
	log.Println(err)

}
