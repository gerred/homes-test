package properties

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"encoding/csv"
)

const fullValuations = `
Date,PropertyID,Address,RatesValuation($)
21/01/2015,1,"4 Link Road, Wellington","$500,000"
21/01/2015,2,"6 Link Road, Wellington","$520,000"
21/01/2015,3,"8 Link Road, Wellington","$540,000"
21/01/2015,4,"10 Link Road, Wellington","$520,000"
21/01/2015,5,"12 Link Road, Wellington","$510,000"
21/01/2014,1,"4 Link Road, Wellington","$480,000"
21/01/2014,2,"6 Link Road, Wellington","$460,000"
21/01/2014,1,"4 Link Road, Wellington","$480,000"
21/01/2015,2,"6 Link Road, Wellington","$550,000"
`

func TestParseBasic(t *testing.T) {
	valuations := `
Date,PropertyID,Address,RatesValuation($)
21/01/2015,1,"4 Link Road, Wellington","$500,000"
`

	csvIn := strings.NewReader(valuations)
	r := csv.NewReader(csvIn)

	properties, err := ParseCSV(r)

	if err != nil {
		fmt.Printf("Got err: %s\n", err)
	}

	date, err := time.Parse(dateFormat, "21/01/2015")
	if err != nil {
		t.Error(err)
	}

	expected := Property{
		Date:      date,
		ID:        1,
		Address:   "4 Link Road, Wellington",
		Valuation: "$500,000",
	}
	actual := (*properties)[1][date]

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Mismatch. Expected: %v, got: %v", expected, actual)
	}
}

func TestParseTopLevelLength(t *testing.T) {

	csvIn := strings.NewReader(fullValuations)
	r := csv.NewReader(csvIn)

	properties, err := ParseCSV(r)

	if err != nil {
		fmt.Printf("Got err: %s\n", err)
	}

	if len(*properties) != 5 {
		fmt.Printf("Expected: 5, actual: %s", len(*properties))
	}
}

func TestOutput(t *testing.T) {
	csvIn := strings.NewReader(fullValuations)
	r := csv.NewReader(csvIn)

	properties, err := ParseCSV(r)

	if err != nil {
		fmt.Printf("Got err: %s\n", err)
	}

	expected := `Date,PropertyID,Address,RatesValuation($)
21/01/2015,4,"10 Link Road, Wellington","$520,000"
21/01/2015,5,"12 Link Road, Wellington","$510,000"
21/01/2015,1,"4 Link Road, Wellington","$500,000"
21/01/2014,1,"4 Link Road, Wellington","$480,000"
21/01/2014,2,"6 Link Road, Wellington","$460,000"
21/01/2015,2,"6 Link Road, Wellington","$550,000"
21/01/2015,3,"8 Link Road, Wellington","$540,000"
`

	if expected != properties.String() {
		fmt.Printf("Mismatch. Expected: %s, got: %s", expected, properties.String())
	}

}
