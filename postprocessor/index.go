package postprocessor

import "github.com/gerred/homes-test/properties"

// Index removes an entry at the modulo of each index
type Index struct {
	Modulo int
}

// Run executes the Duplicate postprocessor
func (ind *Index) Run(p properties.Properties) properties.Properties {
	newProps := properties.Properties{}

	for i := 0; i < len(p); i++ {
		if (i+1)%ind.Modulo == 0 {
			continue
		}

		newProps = append(newProps, p[i])
	}

	return newProps
}
