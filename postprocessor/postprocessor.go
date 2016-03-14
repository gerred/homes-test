package postprocessor

import "github.com/gerred/homes-test/properties"

// Postprocessor is for filters run on an entire list after runtime filters are executed.
type Postprocessor interface {
	Run(properties.Properties) properties.Properties
}
