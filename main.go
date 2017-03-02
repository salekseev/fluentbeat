package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/salekseev/fluentbeat/beater"
)

func main() {
	err := beat.Run("fluentbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
