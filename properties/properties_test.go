package properties

import (
	"reflect"
	"testing"
	"time"
)

var propertiesTable = []struct {
	record           []string
	expectedProperty *Property
}{
	{[]string{"", "", "", "", ""}, nil},
	{[]string{"1", "1 Wanaka Pl", "Wellington", "1/01/15", "300000"}, &Property{1, "1 Wanaka Pl", "Wellington", time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC), 300000}},
}

func TestPropertyParse(t *testing.T) {
	for _, pp := range propertiesTable {
		prop, err := ParseProperty(pp.record)

		if prop == nil && pp.expectedProperty == nil {
			continue
		}
		if err != nil && pp.expectedProperty != nil {
			t.Errorf("%s", err)
		}

		if !reflect.DeepEqual(prop, pp.expectedProperty) {
			t.Errorf("Property parsing.. Expected: %v, got: %v", pp.expectedProperty, prop)
		}
	}
}
