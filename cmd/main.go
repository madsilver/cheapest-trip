package main

import (
	"fmt"
	"os"

	"github.com/silver/cheapest-trip/internal/core/console"
	"github.com/silver/cheapest-trip/internal/core/handler"
	"github.com/silver/cheapest-trip/internal/core/storage"
)

func main() {
	if _, err := storage.CheckCsv(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	go handler.Run()

	console.Run()
}
