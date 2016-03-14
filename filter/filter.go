package filter

import "github.com/gerred/homes-test/properties"

// Filter is implemented to build property filtering middleware.
type Filter interface {
	// Run takes a property and an index of the full list
	Run(*properties.Property) *properties.Property
}
