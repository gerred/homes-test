package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/gerred/homes-test/filter"
	"github.com/gerred/homes-test/pipeline"
	"github.com/gerred/homes-test/postprocessor"
	"github.com/gerred/homes-test/properties"
)

var propsTT = []struct {
	input  [][]string
	output properties.Properties
}{
	{
		input: [][]string{
			[]string{"1", "1 Northburn RD", "WANAKA", "1/01/15", "280000"},
			[]string{"1", "1 Northburn RD", "WANAKA", "1/01/16", "290000"},
			[]string{"2", "1 Mount Ida PL", "WANAKA", "1/01/15", "400000"},
			[]string{"2", "1 Mount Ida PL", "WANAKA", "1/01/15", "800000"},
			[]string{"3", "1 Mount Ida CRES", "WANAKA", "1/01/15", "400000"},
			[]string{"4", "1 Toms Way", "WANAKA", "1/01/15", "400000"},
			[]string{"5", "1 Collins St", "WANAKA", "1/01/15", "400000"},
			[]string{"6", "1 Eely Point Rd", "WANAKA", "1/01/15", "400000"},
			[]string{"7", "1 Baker Gr", "WANAKA", "1/01/15", "700000"},
			[]string{"8", "1 Lakeside Rd", "WANAKA", "1/01/15", "700000"},
			[]string{"9", "1 Islington Dr", "WANAKA", "1/01/15", "700000"},
			[]string{"10", "1 Briar Bank Dr", "WANAKA", "1/01/15", "700000"},
			[]string{"11", "1 Hidden Hills Dr", "WANAKA", "1/01/15", "700000"},
			[]string{"12", "1 Koru Way Dr", "WANAKA", "1/01/15", "700000"},
			[]string{"13", "1 Bovett Dr", "WANAKA", "1/01/15", "700000"},
			[]string{"14", "1 Bovett Pl", "WANAKA", "1/01/15", "700000"},
		},
		output: properties.Properties{
			&properties.Property{ID: 4, Address: "1 Toms Way", Town: "WANAKA", Date: dateBuilder(2015, time.January, 1), Valuation: 400000},
			&properties.Property{ID: 5, Address: "1 Collins St", Town: "WANAKA", Date: dateBuilder(2015, time.January, 1), Valuation: 400000},
			&properties.Property{ID: 6, Address: "1 Eely Point Rd", Town: "WANAKA", Date: dateBuilder(2015, time.January, 1), Valuation: 400000},
			&properties.Property{ID: 7, Address: "1 Baker Gr", Town: "WANAKA", Date: dateBuilder(2015, time.January, 1), Valuation: 700000},
			&properties.Property{ID: 8, Address: "1 Lakeside Rd", Town: "WANAKA", Date: dateBuilder(2015, time.January, 1), Valuation: 700000},
			&properties.Property{ID: 9, Address: "1 Islington Dr", Town: "WANAKA", Date: dateBuilder(2015, time.January, 1), Valuation: 700000},
			&properties.Property{ID: 10, Address: "1 Briar Bank Dr", Town: "WANAKA", Date: dateBuilder(2015, time.January, 1), Valuation: 700000},
			&properties.Property{ID: 11, Address: "1 Hidden Hills Dr", Town: "WANAKA", Date: dateBuilder(2015, time.January, 1), Valuation: 700000},
			&properties.Property{ID: 12, Address: "1 Koru Way Dr", Town: "WANAKA", Date: dateBuilder(2015, time.January, 1), Valuation: 700000},
		}},
}

func TestIntegration(t *testing.T) {
	pipeline := pipeline.DefaultPipeline

	pipeline.RegisterFilter(&filter.CheapFilter{Under: 400000})
	pipeline.RegisterFilter(&filter.SuffixFilter{Suffixes: []string{"AVE", "CRES", "PL"}})

	pipeline.RegisterPostprocessor(&postprocessor.Duplicate{})
	pipeline.RegisterPostprocessor(&postprocessor.Index{Modulo: 10})

	for _, ptt := range propsTT {

		properties, err := pipeline.Run(ptt.input)

		if err != nil {
			t.Errorf("Error: %s", err)
		}

		if len(properties) != len(ptt.output) {
			t.Errorf("Count mismatch. Expected: %d, got: %d", len(ptt.output), len(properties))
		}

		matches := 0
		expectedMatches := 9

		// fmt.Println(ptt.output.String())
		// sort.Sort(properties)
		// fmt.Println(properties.String())

		for _, prop := range properties {
			for _, expectedProp := range ptt.output {
				if reflect.DeepEqual(prop, expectedProp) {
					matches++
				}
			}
		}

		if matches != expectedMatches {
			t.Errorf("Expected %d matches, got %d", expectedMatches, matches)
		}
	}
}

func dateBuilder(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
