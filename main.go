package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/gerred/homes-test/filter"
	"github.com/gerred/homes-test/pipeline"
	"github.com/gerred/homes-test/postprocessor"
)

func main() {
	filename := os.Args[1]

	in, err := os.Open(filename)
	defer in.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	csvFile := csv.NewReader(in)
	data, err := csvFile.ReadAll()

	dataNoHeader := data[1:]

	if err != nil {
		fmt.Println(err)
		return
	}

	pipeline := pipeline.DefaultPipeline

	pipeline.RegisterFilter(&filter.CheapFilter{Under: 700000})
	pipeline.RegisterFilter(&filter.SuffixFilter{Suffixes: []string{"AVE", "CRES", "PL"}})

	pipeline.RegisterPostprocessor(&postprocessor.Duplicate{})
	pipeline.RegisterPostprocessor(&postprocessor.Index{Modulo: 10})

	properties, err := pipeline.Run(dataNoHeader)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(properties.String())
}
