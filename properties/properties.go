package properties

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
	"time"
)

const dateFormat = "02/01/2006"

// Property represents the
type Property struct {
	date      time.Time
	id        int
	address   string
	valuation string
}

// Properties represents a collection of properties and their valuation over time
type Properties map[int]map[time.Time]Property

// String outputs a pretty printed table of properties
func (p *Properties) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Date,PropertyID,Address,RatesValuation($)\n")
	for _, propertyMap := range *p {
		for _, property := range propertyMap {
			buffer.WriteString(fmt.Sprintf("%s,%d,\"%s\",\"%s\"\n", property.date.Format(dateFormat), property.id, property.address, property.valuation))
		}
	}

	return buffer.String()
}

// ParseCSV takes an arbitrary reader of CSV-formatted data and returns Properties
func ParseCSV(reader *csv.Reader) (*Properties, error) {
	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	properties := Properties{}

	for i, record := range records {
		if i == 0 {
			continue
		}

		property, err := parseProperty(record)
		if err != nil {
			return nil, err
		}

		if _, ok := properties[property.id]; !ok {
			properties[property.id] = map[time.Time]Property{}
		}

		properties[property.id][property.date] = *property
	}

	return &properties, nil
}

// parseProperty returns a single Property from a slice record
func parseProperty(record []string) (*Property, error) {
	date, err := time.Parse(dateFormat, record[0])
	if err != nil {
		return nil, err
	}

	id, err := strconv.ParseInt(record[1], 10, 8)
	if err != nil {
		return nil, err
	}
	propertyID := int(id)

	return &Property{date, propertyID, record[2], record[3]}, nil
}
