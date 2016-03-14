package postprocessor

import (
	"testing"
	"time"

	"github.com/gerred/homes-test/properties"
)

var duplicatePP = &Duplicate{}

var duplicatesTable = []struct {
	properties     []*properties.Property
	expectedLength int
}{
	{[]*properties.Property{}, 0},
	{[]*properties.Property{
		{ID: 1, Date: dateBuilder(2015, time.January, 1)},
		{ID: 1, Date: dateBuilder(2015, time.February, 1)},
	}, 2},
	{[]*properties.Property{
		{ID: 1, Date: dateBuilder(2015, time.January, 1)},
		{ID: 1, Date: dateBuilder(2016, time.January, 1)},
	}, 2},
	{[]*properties.Property{
		{ID: 1, Date: dateBuilder(2015, time.January, 1)},
		{ID: 2, Date: dateBuilder(2015, time.January, 1)},
	}, 2},
	{[]*properties.Property{
		{ID: 1, Date: dateBuilder(2015, time.January, 1)},
		{ID: 1, Date: dateBuilder(2015, time.February, 1)},
		{ID: 1, Date: dateBuilder(2015, time.January, 1)},
	}, 1},
}

func TestRemovesDuplicates(t *testing.T) {
	for _, pp := range duplicatesTable {
		properties := duplicatePP.Run(pp.properties)
		if len(properties) != pp.expectedLength {
			t.Errorf("Duplicates postprocessor. Expected length: %d, got: %d", pp.expectedLength, len(properties))
		}
	}
}

func dateBuilder(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
