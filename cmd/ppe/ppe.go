package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilario-pierbattista/parallel-phpunit-executor/internal/balancer"
)

func main() {
	var testPath string

	flag.StringVar(&testPath, "path", "", "Test path")

	flag.Parse()

	if testPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	chunks, err := balancer.MakeChunks(testPath, 30)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(chunks)

	fmt.Println("Number of chunks:", len(chunks))
}
