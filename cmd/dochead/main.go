package main

import (
	"os"
	"errors"
	"fmt"
    "github.com/bitionaire/dochead"
)

func main() {
	files := os.Args[1:]
	if len(files) != 1 {
		error := errors.New("what's wrong with ya?")
		panic(error)
	} else {
		file := files[0]
		
        apiDefinition := dochead.ReadAPIDefinition(file)
        fmt.Printf("Api %s\n", apiDefinition)
	}
}
