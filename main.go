package main

import (
	"Designweb/config"
	"log"
	"net/http"
)

func main() {
	var Start string = "########################################################\n################\t FUDEMY\t################\n################################################\n################\tBURAK ATAŞ\t################\n########################################################\n########################################################\n"

	log.Println(Start)

	log.Printf("Starting Server -->>%v \n\n", 8080)

	log.Fatal(http.ListenAndServe(":8080", config.Router_Config()))
}
