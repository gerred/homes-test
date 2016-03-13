package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/gerred/homes-test/properties"
)

func main() {
	filename := os.Args[1]

	inFile, err := os.Open(filename)
	defer inFile.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	properties, err := properties.ParseCSV(csv.NewReader(inFile))

	fmt.Println(properties)
}
