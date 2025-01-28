package main

import (
	// "fmt"
	"log"
	"realmrovers/cmd"
	// "realmrovers/config"
	// "realmrovers/db"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}

}
