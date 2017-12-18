package main

import (
	"log"

	"github.com/dgx80/goio"
	"github.com/fatih/color"
)

var path = "./Makefile"

func main() {

	if goio.FileExist(path) {
		log.Println("file exist")
		failed("file exist")
	}

	if !goio.CreateFile(path) {
		failed("unable to create file")
	}
	writeDefaultsFile()
}

func failed(message string) {
	color.Red(message)
	log.Fatalln("teminate with an error")
}

func writeDefaultsFile() {
	var lines = []string{
		"#administred by github.com/dgx80/maker\n",
		".PHONY: help\n\n",
		"help: ## help - Display callable targets. Hint: you can write just make to see help\n\t",
		`@grep -E '(^[a-zA-Z_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-20s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'`,
		"\n",
	}

	for _, line := range lines {
		goio.WriteFile(path, line)
	}
}
