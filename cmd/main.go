package main

import (
	"fmt"
	"log"

	"github.com/kazufusa/comlist"
)

func main() {
	coms, err := comlist.Coms()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coms)
}
