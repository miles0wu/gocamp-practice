package main

import (
	"log"
)

func main() {
	query := "SELECT 1"
	err := Dao(query)
	if IsNoRow(err) {
		log.Printf("%+v\n", err)
		return
	}
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}
