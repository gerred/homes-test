package filter

import (
	"testing"

	"github.com/gerred/homes-test/properties"
)

var cheapFilter = &CheapFilter{
	Under: 300000,
}

var propertiesTable = []struct {
	property *properties.Property
	filtered bool
}{
	{nil, true},
	{&properties.Property{Valuation: 500000}, false},
	{&properties.Property{Valuation: 100000}, true},
	{&properties.Property{Valuation: 300000}, false},
}

func TestRemovesCheap(t *testing.T) {
	for _, pp := range propertiesTable {
		prop := cheapFilter.Run(pp.property)
		expected := prop == nil
		if expected != pp.filtered {
			t.Errorf("Cheap filtered. Expected: %t, want: %t", expected, pp.filtered)
		}
	}
}
