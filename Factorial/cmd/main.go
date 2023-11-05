package main

import (
	"flag"
	"go/course/factorial/pkg/calculations"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func main() {
	var loggingFlag = flag.Bool("log", false, "if flag is set, you will see logging msgs from inner functions")
	flag.Parse()

	log.SetLevel(log.InfoLevel)
	log.SetOutput(os.Stdout)

	if len(os.Args) < 2 {
		log.Fatal("args should contain minimum 1 parameter")
	}
	number, err := strconv.ParseInt(flag.Arg(0), 10, 64)
	if err != nil {
		log.Fatal("args should contain number for factorial calculation")
	}

	result, err := calculations.CalculateFactorial(number, *loggingFlag)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Factorial of %d is %d", number, result)
}
