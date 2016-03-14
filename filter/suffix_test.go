package filter

import (
	"testing"

	"github.com/gerred/homes-test/properties"
)

var suffixFilter = &SuffixFilter{
	Suffixes: []string{"AVE", "CRES"},
}

var suffixesTable = []struct {
	property *properties.Property
	filtered bool
}{
	{nil, true},
	{&properties.Property{Address: "51 Wanaka Pl"}, false},
	{&properties.Property{Address: "10 Test AVE"}, true},
	{&properties.Property{Address: "10 Test Cres"}, true},
}

func TestRemovesSuffix(t *testing.T) {
	for _, pp := range suffixesTable {
		prop := suffixFilter.Run(pp.property)
		expected := prop == nil
		if expected != pp.filtered {
			t.Errorf("Suffix filtered. Expected: %t, want: %t", expected, pp.filtered)
		}
	}
}
