package main

import (
	"log"

	"github.com/CarosDrean/utils/cmd/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
