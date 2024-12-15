package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	interpreter "github.com/0xmukesh/brainfreeze/internal"
)

func main() {
	var filename string
	flag.StringVar(&filename, "file", "", "brainf*ck filename")

	flag.Parse()

	if filename == "" {
		fmt.Println("missing file")
		os.Exit(1)
	}

	ext := path.Ext(filename)
	if ext != ".bf" {
		fmt.Println("invalid file extension")
		os.Exit(1)
	}

	input, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	iptr := interpreter.NewInterpreter(input)
	iptr.Run()
}
