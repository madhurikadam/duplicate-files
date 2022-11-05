package main

import (
	"fmt"
	"log"
	"os"

	"github.com/madhurikadam/duplicate-files/internal/service"
)

func main() {

	dirName := os.Args[1]
	duplicateSvs := service.New()
	duplicate, err := duplicateSvs.GetDuplicates(dirName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(duplicate)
}
