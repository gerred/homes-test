package filter

import "github.com/gerred/homes-test/properties"

// CheapFilter filters properties under a certain valuation
type CheapFilter struct {
	Under int64
}

// Run executes the CheapFilter
func (c *CheapFilter) Run(p *properties.Property) *properties.Property {
	if p == nil || p.Valuation < c.Under {
		return nil
	}
	return p
}
