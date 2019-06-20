package main

import (
	"fmt"
	"os"

	"github.com/mercatormaps/go-mapcss"
)

func main() {
	f, err := os.Open("testdata/sample.mapcss")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ss, err := mapcss.NewStylesheet(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := ss.Parse(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
