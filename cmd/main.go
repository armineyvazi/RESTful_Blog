package main

import (
	"fmt"
	"log"

	"restful_blog/cmd/api"
)

func main() {
	if err := api.Run(); err != nil {
		log.Fatalf("Error %v", err)
		//os.Exit(1)
	}

	fmt.Println("Application exited successfully ... :))")

}
