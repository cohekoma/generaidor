package main

import (
	"log"

	"github.com/cohekoma/generaidor/gen"
)

func main() {

	err := gen.HandleStaticGeneration()

	if err != nil {
		log.Fatal("Fail to generate static content, please try again.")
	}
}
