package postprocessor

import (
	"testing"

	"github.com/gerred/homes-test/properties"
)

var indexPP = &Index{Modulo: 3}

var indexTable = []struct {
	properties  []*properties.Property
	expectedIDs []int
}{
	{
		[]*properties.Property{
			{ID: 1},
			{ID: 2},
			{ID: 3},
			{ID: 4},
		},
		[]int{1, 2, 4},
	},
}

func TestRemovesIndex(t *testing.T) {
	for _, pp := range indexTable {
		properties := indexPP.Run(pp.properties)
		for _, p := range properties {
			if !contains(pp.expectedIDs, p.ID) {
				t.Errorf("Expected %d to be filtered.", p.ID)
			}
		}

	}
}
func contains(c []int, i int) bool {
	for _, id := range c {
		if i == id {
			return true
		}
	}
	return false
}
