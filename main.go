package main

import (
	"log"
	"net/http"
	"new_CodingTime/config"
)

func main() {
	PORT := 8080
	var Start string = "########################################################\n################\t FUDEMY\t################\n################################################\n################\tBURAK ATAŞ\t################\n########################################################\n########################################################\n"
	log.Println(Start)

	log.Printf("Starting Server -->>%v \n\n", PORT)
	log.Fatal(http.ListenAndServe(":8000", config.Router_Config()))
}
