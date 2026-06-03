package main

import (
	"fmt"
	"log"

	"github.com/Mate2xo/gator/internal/config"
)

func main() {
	file, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", file)

	err = config.SetUser("mate")
	if err != nil {
		log.Fatal(err)
	}

	file, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", file)
}
