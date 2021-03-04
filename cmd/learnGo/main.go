package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/DeepikaAzad/LearnGo"
)

func main() {
	file := flag.String("file", "gopher.json", "The Json file with CYON adventure story.")
	flag.Parse()

	fmt.Printf("Using the file %s\n", *file)
	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	var story LearnGo.Story
}
