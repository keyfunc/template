package main

import (
	"log"

	"github.com/promonkeyli/goas/pkg/goas"
)

func main() {

	config := goas.Config{
		Dirs:   []string{"./cmd/server", "./internal"},
		Output: ".docs/api",
	}

	if err := goas.Run(config); err != nil {
		log.Fatalf("Failed to generate OpenAPI: %v", err)
	}

}
