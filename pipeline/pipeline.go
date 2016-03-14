package pipeline

import (
	"sort"
	"sync"

	"github.com/gerred/homes-test/filter"
	"github.com/gerred/homes-test/postprocessor"
	"github.com/gerred/homes-test/properties"
)

// Pipeline is the full lifecycle of parsing a set of properties, and applying filters to that list in chunks.
type Pipeline struct {
	ParseFunc        func([]string) (*properties.Property, error)
	postprocessors   []postprocessor.Postprocessor
	IgnoreParseError bool
	properties       properties.Properties
	filters          []filter.Filter
	ChunkSize        int
}

// DefaultPipeline is an empty pipeline, set to parse properties with a chunk size of 10.
var DefaultPipeline = &Pipeline{
	ParseFunc:        properties.ParseProperty,
	IgnoreParseError: true,
	properties:       []*properties.Property{},
	postprocessors:   []postprocessor.Postprocessor{},
	filters:          []filter.Filter{},
	ChunkSize:        10,
}

// RegisterFilter appends a new filter into the pipeline. Filters are run as chunks at runtime
func (p *Pipeline) RegisterFilter(filter filter.Filter) {
	p.filters = append(p.filters, filter)
}

// RegisterPostprocessor adds a postprocessor into the pipeline. Postprocessors are run on the entire list after filters have been run.
func (p *Pipeline) RegisterPostprocessor(prc postprocessor.Postprocessor) {
	p.postprocessors = append(p.postprocessors, prc)
}

// Run executes a pipeline concurrently given an input of property string slices. This function assumes a header has already been stripped, and is not intended to do CSV parsing.
func (p *Pipeline) Run(input [][]string) (properties.Properties, error) {
	for _, record := range input {
		property, err := p.ParseFunc(record)

		if err != nil {
			if p.IgnoreParseError == false {
				return nil, err
			}
		} else {
			p.properties = append(p.properties, property)
		}
	}

	p.properties = runFilterChain(p.filters, p.properties, p.ChunkSize)
	p.properties = runPostprocessorChain(p.postprocessors, p.properties)

	return p.properties, nil
}

func runFilterChain(filters []filter.Filter, p properties.Properties, chunkSize int) properties.Properties {
	for _, filter := range filters {
		pChan := make(chan *properties.Property, len(p))
		var wg sync.WaitGroup

		var filterChunk = func(c []*properties.Property) {
			for _, property := range c {
				val := filter.Run(property)
				if val != nil {
					pChan <- property
				}
			}
			wg.Done()
		}

		chunk := []*properties.Property{}
		for i := 0; i < len(p); i++ {
			chunk = append(chunk, p[i])
			if len(chunk) == chunkSize {
				wg.Add(1)
				go filterChunk(chunk)
				chunk = []*properties.Property{}
			}
		}
		wg.Add(1)
		go filterChunk(chunk)
		wg.Wait()

		close(pChan)
		p = []*properties.Property{}

	appendProperties:
		for {
			prop, more := <-pChan
			if !more {
				break appendProperties
			}
			p = append(p, prop)
		}

	}
	return p
}

func runPostprocessorChain(pp []postprocessor.Postprocessor, p properties.Properties) properties.Properties {
	for _, postproc := range pp {
		p = postproc.Run(p)
		sort.Sort(p)
	}

	return p
}
