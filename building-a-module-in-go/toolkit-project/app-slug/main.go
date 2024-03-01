package main

import (
	"log"

	"github.com/arturbaccarin/toolkit"
)

func main() {
	toSlug := "NOW is the Time 123"

	var tools toolkit.Tools

	slugified, err := tools.Slugify(toSlug)
	if err != nil {
		log.Println(err)
	}

	log.Println(slugified)
}
