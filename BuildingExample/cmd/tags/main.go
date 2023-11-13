package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.ErrorLevel)
}

func main() {
	a := 1
	b := 2
	log.Error("some error happened")
	sum := a + b
	log.Infof("sum successfully evaluated %d", sum)
	fmt.Printf("sum is: %d", sum)
}
