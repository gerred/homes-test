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

type RemoveIDFilter struct {
	ID int
}

func (r *RemoveIDFilter) Run(p *properties.Property) *properties.Property {
	if p.ID == r.ID {
		return nil
	}
	return p
}

// Successful criteria: Spy runs on all elements (filter runs on all), even if the filtered length is lower
func TestFilterChainSpyFirst(t *testing.T) {
	spyFilter := &SpyFilter{}
	removeIDFilter := &RemoveIDFilter{ID: 3}

	props := properties.Properties{
		&properties.Property{},
		&properties.Property{},
		&properties.Property{ID: 3},
		&properties.Property{},
		&properties.Property{},
		&properties.Property{},
		&properties.Property{},
		&properties.Property{},
		&properties.Property{},
		&properties.Property{},
	}

	newProps := runFilterChain([]filter.Filter{spyFilter, removeIDFilter}, props, 10)
	expectedFinalLength := len(props) - 1

	if len(newProps) != expectedFinalLength {
		t.Errorf("Pipeline length mismatch. Got: %d, expected: %d", len(newProps), expectedFinalLength)
	}

	if int(spyFilter.callCount) != len(props) {
		t.Errorf("Not enough filters got called. Expected: %d, saw: %d", len(newProps), spyFilter.callCount)
	}

}

func TestFilterChainSpyLast(t *testing.T) {
	spyFilter := &SpyFilter{}
	removeIDFilter := &RemoveIDFilter{ID: 4}

	props := properties.Properties{
		&properties.Property{},
		&properties.Property{},
		&properties.Property{ID: 4},
		&properties.Property{},
		&properties.Property{},
		&properties.Property{},
		&properties.Property{},
		&properties.Property{},
		&properties.Property{},
		&properties.Property{},
	}

	newProps := runFilterChain([]filter.Filter{removeIDFilter, spyFilter}, props, 10)

	expectedFinalLength := len(props) - 1

	if len(newProps) != expectedFinalLength {
		t.Errorf("Pipeline length mismatch. Got: %d, expected: %d", len(newProps), expectedFinalLength)
	}

	if int(spyFilter.callCount) != expectedFinalLength {
		t.Errorf("Filter call mismatch. Expected: %d, saw: %d", spyFilter.callCount, expectedFinalLength)
	}

}
