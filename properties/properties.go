package properties

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

const DateFormat = "2/01/06"

// Property represents a single property and it's valuation
type Property struct {
	ID        int
	Address   string
	Town      string
	Date      time.Time
	Valuation int64
}

func (p *Property) String() string {
	return fmt.Sprintf("%d,%s,%s,%s,%d", p.ID, p.Address, p.Town, p.Date, p.Valuation)
}

// Properties represents a collection of properties and their valuation over time
type Properties []*Property

func (p *Properties) String() string {
	var buffer bytes.Buffer

	for _, r := range *p {
		buffer.WriteString(fmt.Sprintf("%s\n", r))
	}

	return buffer.String()
}

// ParseProperty converts a string slice into a fully qualified Property
func ParseProperty(record []string) (*Property, error) {
	date, err := time.Parse(DateFormat, record[3])
	if err != nil {
		return nil, err
	}

	id, err := strconv.ParseInt(record[0], 10, 8)
	if err != nil {
		return nil, err
	}
	propertyID := int(id)

	v, err := strconv.ParseInt(record[4], 10, 64)
	if err != nil {
		return nil, err
	}

	return &Property{propertyID, record[1], record[2], date, v}, nil
}
