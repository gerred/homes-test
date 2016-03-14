package pipeline

import (
	"sync/atomic"
	"testing"

	"github.com/gerred/homes-test/filter"
	"github.com/gerred/homes-test/properties"
)

type SpyFilter struct {
	callCount uint64
}

func (s *SpyFilter) Run(p *properties.Property) *properties.Property {
	atomic.AddUint64(&s.callCount, 1)

	return p
}

func TestFilterChain(t *testing.T) {
	spyFilter := &SpyFilter{}

	props := properties.Properties{
		&properties.Property{},
		&properties.Property{},
		&properties.Property{},
		&properties.Property{},
		&properties.Property{},
		&properties.Property{},
		&properties.Property{},
		&properties.Property{},
		&properties.Property{},
		&properties.Property{},
	}

	newProps := runFilterChain([]filter.Filter{spyFilter}, props, 10)

	if len(newProps) != len(props) {
		t.Errorf("Pipeline length mismatch. Got: %d, expected: %d", len(newProps), len(props))
	}

	if int(spyFilter.callCount) != len(props) {
		t.Errorf("Not enough filters got called. Expected: %d, saw: %d", len(props), spyFilter.callCount)
	}

}
