package main

import (
	"fmt"
	"go/course/files/internal/yaml"
	"os"

	"github.com/PonomarevAlexxander/files"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.WarnLevel)

	if len(os.Args) < 2 {
		log.Fatal("args should contain path to .yml file")
	}

	ymlPath := os.Args[1]
	ymlData, err := os.ReadFile(ymlPath)
	if err != nil {
		log.Fatal(err)
	}

	filesList, err := yaml.Unmarshal(ymlData)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range filesList {
		contains, err := files.Contains(file.Path, file.Substring)
		if err != nil {
			log.Warn(err)
			continue
		}
		var result string
		if !contains {
			result = "no"
		}
		fmt.Printf("File '%s' contains %s substring: '%s'\n", file.Path, result, file.Substring)
	}
}
