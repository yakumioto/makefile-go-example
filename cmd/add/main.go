package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/yakumioto/go-makefile-example/internal/utils"
)

func main() {
	if len(os.Args) != 3 {
		log.Panic("Requires two int parameters")
	}

	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Panic("Parameter 1 needs to be of type int")
	}

	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Panic("Parameter 2 needs to be of type int")
	}

	fmt.Printf("%v + %v = %v\n", a, b, utils.Add(a, b))
}
