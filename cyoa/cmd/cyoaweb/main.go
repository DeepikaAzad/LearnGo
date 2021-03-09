package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/DeepikaAzad/LearnGo/cyoa"
)

func main() {
	file := flag.String("file", "gopher.json", "The Json file with CYON adventure story.")
	flag.Parse()

	fmt.Printf("Using the file %s\n", *file)
	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JSONStory(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", story)
}
