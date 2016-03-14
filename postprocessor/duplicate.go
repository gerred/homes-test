package postprocessor

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/gerred/homes-test/properties"
)

// Duplicate removes an entry from a list entirely if it has any duplicates
type Duplicate struct {
}

// Run executes the Duplicate postprocessor
func (d *Duplicate) Run(p properties.Properties) properties.Properties {
	props := map[string]properties.Properties{}

	for _, prop := range p {
		idStr := string(prop.ID)
		dateStr := prop.Date.Format(properties.DateFormat)

		hash := md5.New()
		hash.Write([]byte(idStr + dateStr))
		hashStr := hex.EncodeToString(hash.Sum(nil))

		if _, ok := props[hashStr]; !ok {
			props[hashStr] = properties.Properties{prop}
			continue
		}
		props[hashStr] = append(props[hashStr], prop)

	}

	newProps := properties.Properties{}

	for _, v := range props {
		if len(v) == 1 {
			newProps = append(newProps, v[0])
		}
	}

	return newProps
}
