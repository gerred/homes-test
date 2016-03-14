package filter

import (
	"strings"

	"github.com/gerred/homes-test/properties"
)

// SuffixFilter filters properties that have a particular suffix
type SuffixFilter struct {
	Suffixes []string
}

// Run executes the CheapFilter
func (s *SuffixFilter) Run(p *properties.Property) *properties.Property {
	splitAddress := strings.Split(p.Address, " ")
	suffix := splitAddress[len(splitAddress)-1]

	for _, s := range s.Suffixes {
		if strings.ToLower(s) == strings.ToLower(suffix) {
			return nil
		}
	}
	return p

}
