package main

import (
	"fmt"
	"log"
)

func main() {
	log.Print("test")
	err := fmt.Errorf("err")
	log.Fatal(err)
}
