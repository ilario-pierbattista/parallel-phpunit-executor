package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ilario-pierbattista/parallel-phpunit-executor/internal/balancer"
	"github.com/ilario-pierbattista/parallel-phpunit-executor/internal/phpunitconf"
)

func main() {
	var testPath string
	var phpunitConfig string
	var debug bool

	flag.StringVar(&testPath, "path", "", "Test path")
	flag.StringVar(&phpunitConfig, "config", "", "PHPUnit configuration file")
	flag.BoolVar(&debug, "debug", false, "Show debug informations")

	flag.Parse()

	if testPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	if phpunitConfig == "" {
		flag.Usage()
		os.Exit(1)
	}

	if debug {
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Println(*f)
		})
	}

	chunks, err := balancer.MakeChunks(testPath, 30)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Println(chunks)

	fmt.Println("Number of chunks:", len(chunks))

	_, err = phpunitconf.Parse(strings.TrimLeft(phpunitConfig, "="))
	fmt.Println(err)
}
